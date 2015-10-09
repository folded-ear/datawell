package main

import (
	"github.com/folded-ear/datawell/model"
	"log"
)

type EditUserCommand struct {
	Command
	name     *string
	username *string
	id       *int64
}

func (c *EditUserCommand) resetFlags() {
	if *c.id == 0 {
		c.id = nil
	}
	if *c.username == "" {
		c.username = nil
	}
}

func (c *EditUserCommand) checkFlags() {
	if (c.id != nil && c.username != nil) || (c.id == nil && c.username == nil) {
		log.Fatalf("Exactly one of --id and --username must be provided")
	}
	if *c.name == "" {
		log.Fatalf("You have to specify something to update.")
	}
}

var editUserCmd = &EditUserCommand{
	Command: Command{
		Name:    "edit-user",
		Usage:   "",
		Summary: "edit a user in the system",
	},
}

func init() {
	editUserCmd.Run = editUserRun
	editUserCmd.name = editUserCmd.Flag.String("name", "", "The user's updated name")
	editUserCmd.username = editUserCmd.Flag.String("username", "", "The username of the user to update (disallowed if --id is provided)")
	editUserCmd.id = editUserCmd.Flag.Int64("id", 0, "The id of the user to update (disallowed if --username is provided)")
}

func editUserRun(cmd *Command, args ...string) {
	editUserCmd.resetFlags()
	editUserCmd.checkFlags()
	db, err := model.Gorm()
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	user := &model.User{}
	if editUserCmd.id != nil {
		user.ID = *editUserCmd.id
	} else {
		user.Username = *editUserCmd.username
	}
	err = db.Where(user).First(user).Error
	if err != nil {
		log.Fatal(err)
	}
	user.Name = *editUserCmd.name
	db.Save(user)
}
