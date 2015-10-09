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

var (
	editUserCmd = &Command{
		Name:    "edit-user",
		Usage:   "",
		Summary: "edit a user in the system",
		Run:     editUserRun,
	}
	eu_id       *int64
	eu_username *string
	eu_name     *string
)

func init() {
	eu_name = editUserCmd.Flag.String("name", "", "The user's updated name")
	eu_username = editUserCmd.Flag.String("username", "", "The username of the user to update (disallowed if --id is provided)")
	eu_id = editUserCmd.Flag.Int64("id", 0, "The id of the user to update (disallowed if --username is provided)")
}

func editUserRun(cmd *Command, args ...string) {
	if *eu_id == 0 {
		eu_id = nil
	}
	if *eu_username == "" {
		eu_username = nil
	}
	if (eu_id != nil && eu_username != nil) || (eu_id == nil && eu_username == nil) {
		log.Fatalf("Exactly one of --id and --username must be provided")
	}
	if *eu_name == "" {
		log.Fatalf("You have to specify something to update.")
	}
	db, err := model.Gorm()
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	user := &model.User{}
	if eu_id != nil {
		user.ID = *eu_id
	} else {
		user.Username = *eu_username
	}
	err = db.Where(user).First(user).Error
	if err != nil {
		log.Fatal(err)
	}
	user.Name = *eu_name
	db.Save(user)
}
