package main

import (
	"github.com/folded-ear/datawell/model"
	"log"
)

type UserEditCommand struct {
	Command
	name     *string
	username *string
	id       *int64
}

func (c *UserEditCommand) resetFlags() {
	if *c.id == 0 {
		c.id = nil
	}
	if *c.username == "" {
		c.username = nil
	}
}

func (c *UserEditCommand) checkFlags() {
	if (c.id != nil && c.username != nil) || (c.id == nil && c.username == nil) {
		log.Fatalf("Exactly one of --id and --username must be provided")
	}
	if *c.name == "" {
		log.Fatalf("You have to specify something to update.")
	}
}

var userEditCmd = &UserEditCommand{
	Command: Command{
		Name:    "user-edit",
		Usage:   "",
		Summary: "edit a user in the system",
	},
}

func init() {
	userEditCmd.Run = userEditRun
	userEditCmd.name = userEditCmd.Flag.String("name", "", "The user's updated name")
	userEditCmd.username = userEditCmd.Flag.String("username", "", "The username of the user to update (disallowed if --id is provided)")
	userEditCmd.id = userEditCmd.Flag.Int64("id", 0, "The id of the user to update (disallowed if --username is provided)")
}

func userEditRun(cmd *Command, args ...string) {
	userEditCmd.resetFlags()
	userEditCmd.checkFlags()
	db, err := model.Gorm()
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	user := &model.User{}
	if userEditCmd.id != nil {
		user.ID = *userEditCmd.id
	} else {
		user.Username = *userEditCmd.username
	}
	err = db.Where(user).First(user).Error
	if err != nil {
		log.Fatal(err)
	}
	user.Name = *userEditCmd.name
	db.Save(user)
}
