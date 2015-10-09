package main

import (
	"fmt"
	"github.com/folded-ear/datawell/model"
	"log"
	"strconv"
)

var (
	listUsersCmd = &Command{
		Name:    "list-users",
		Usage:   "",
		Summary: "list all users in the system",
		Run:     listUsersRun,
	}
)

func listUsersRun(cmd *Command, args ...string) {
	users := []model.User{}
	db, err := model.Gorm()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Order("name").Order("username").Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}
	idLen, nameLen, unLen := 2, 4, 8
	for _, u := range users {
		l := len(strconv.FormatInt(u.ID, 10))
		if l > idLen {
			idLen = l
		}
		l = len(u.Name)
		if l > nameLen {
			nameLen = l
		}
		l = len(u.Username)
		if l > unLen {
			unLen = l
		}
	}
	fmt.Printf("| %*s | %-*s | %-*s |\n", idLen, "id", unLen, "username", nameLen, "name")
	for _, u := range users {
		fmt.Printf("| %*d | %-*s | %-*s |\n", idLen, u.ID, unLen, u.Username, nameLen, u.Name)
	}
	fmt.Printf("%d user(s)\n", len(users))
}
