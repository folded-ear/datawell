/*
Package main is the main entrypoint to the datawell application, including
configuration loading.
*/
package main

import (
	"database/sql"
	"fmt"
	"github.com/folded-ear/datawell/model"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	config, err := newConfig()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		return
	}

	fmt.Printf("hello, world!  here's a gorm error: %v.  And here's a config flag: %v\n", gorm.InvalidSql, config.env)

	fmt.Printf("driver: %v, open: %v\n", config.driverName, config.dataSourceName)

	db, err := sql.Open(config.driverName, config.dataSourceName)
	if err != nil {
		fmt.Printf("connect error: %v\n", err)
		return
	}
	var now time.Time
	err = db.QueryRow("select now()").Scan(&now)
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}
	fmt.Printf("it's %v\n", now)

	gorm, err := gorm.Open(config.driverName, db)
	if err != nil {
		fmt.Printf("gorm open error: %v\n", err)
		return
	}

	tx := gorm.Begin()

	//	user := model.User{
	//		Name:     "Barney Boisvert",
	//		Username: "barneyb"}
	//	tx.Create(&user)

	users := make([]model.User, 10)
	tx.Find(&users)
	for _, u := range users {
		fmt.Printf("user #%v: %v\n", u.ID, u.Username)
	}

	tx.Commit()
}
