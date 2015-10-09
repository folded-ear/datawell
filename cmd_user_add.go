package main

import (
	"fmt"
	"github.com/folded-ear/datawell/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

type UserAddCommand struct {
	Command
	name     *string
	username *string
}

const (
	InitialPasswordLength = 12
	PasswordHashCost      = 13 // takes about 700ms on my laptop
)

var (
	userAddCmd = &UserAddCommand{
		Command: Command{
			Name:    "user-add",
			Usage:   "",
			Summary: "add a new user to the system",
		},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
	userAddCmd.Run = userAddRun
	userAddCmd.name = userAddCmd.Flag.String("name", "", "The name of the new user")
	userAddCmd.username = userAddCmd.Flag.String("username", "", "The username of the new user (must be unique)")
}

func createPassword() string {
	password := []rune{}
	var offset, count int
	for i := 0; i < InitialPasswordLength; i++ {
		if i < 4 {
			// start with uppercase
			offset = 65
			count = 26
		} else if i > InitialPasswordLength-4 {
			// end with numbers
			offset = 48
			count = 10
		} else {
			// middle is lowercase
			offset = 97
			count = 26
		}
		password = append(password, rune(offset+rand.Intn(count)))
	}
	return string(password)
}

func userAddRun(cmd *Command, args ...string) {
	if *userAddCmd.username == "" {
		log.Fatal("A username must be specified")
	}
	passwd := createPassword()
	encPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), PasswordHashCost)
	if err != nil {
		log.Fatalf("Failed to bcrypt password: %v", err)
	}
	user := model.User{
		Username: *userAddCmd.username,
		Name:     *userAddCmd.name,
		Passhash: string(encPasswd),
	}
	db, err := model.Gorm()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created user '%s' (#%d) with initial password '%s'\n", user.Username, user.ID, passwd)
}
