package main

import (
	"github.com/folded-ear/datawell/model"
	"log"
)

type UserDeleteCommand struct {
	Command
	username *string
	id       *int64
}

func (c *UserDeleteCommand) resetFlags() {
	if *c.id == 0 {
		c.id = nil
	}
	if *c.username == "" {
		c.username = nil
	}
}

func (c *UserDeleteCommand) checkFlags() {
	if (c.id != nil && c.username != nil) || (c.id == nil && c.username == nil) {
		log.Fatalf("Exactly one of --id and --username must be provided")
	}
}

var userDeleteCmd = &UserDeleteCommand{
	Command: Command{
		Name:    "user-delete",
		Usage:   "",
		Summary: "delete a user from the system",
	},
}

func init() {
	userDeleteCmd.Run = userDeleteRun
	userDeleteCmd.username = userDeleteCmd.Flag.String("username", "", "The username of the user to delete (disallowed if --id is provided)")
	userDeleteCmd.id = userDeleteCmd.Flag.Int64("id", 0, "The id of the user to delete (disallowed if --username is provided)")
}

func userDeleteRun(cmd *Command, args ...string) {
	userDeleteCmd.resetFlags()
	userDeleteCmd.checkFlags()
	db, err := model.Gorm()
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	user := &model.User{}
	if userDeleteCmd.id != nil {
		user.ID = *userDeleteCmd.id
	} else {
		user.Username = *userDeleteCmd.username
	}
	err = db.Where(user).First(user).Error
	if err != nil {
		log.Fatal(err)
	}
	db.Delete(user)
}
