package main

import (
	"fmt"
	"github.com/folded-ear/datawell/config"
	"github.com/folded-ear/datawell/model"
	"log"
	"time"
)

var demoCmd = &Command{
	Name:    "demo",
	Usage:   "i do a little demo schmanky",
	Summary: "run the demo",
	Run:     demoRun,
}

func demoRun(cmd *Command, args ...string) {
	config := config.LoadConfig()

	fmt.Printf("driver: %v, open: %v\n", config.DriverName, config.DataSourceName)

	db, err := model.DB()
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	var now time.Time
	err = db.QueryRow("select now()").Scan(&now)
	if err != nil {
		log.Fatalf("query error: %v\n", err)
	}
	fmt.Printf("it's %v\n", now)

	gorm, err := model.GORM()
	if err != nil {
		log.Fatalf("gorm open error: %v\n", err)
	}
	gorm.LogMode(true)

	tx := gorm.Begin()

	user := model.User{}
	fmt.Println(tx.FirstOrCreate(&user, model.User{
		Name:     "Barney Boisvert",
		Username: "barneyb"}).Error)

	coffee := model.Tag{}
	fmt.Println(tx.FirstOrCreate(&coffee, model.Tag{
		UserID: user.ID,
		Tag:    "coffee"}).Error)
	desk := model.Tag{}
	fmt.Println(tx.FirstOrCreate(&desk, model.Tag{
		UserID: user.ID,
		Tag:    "desk"}).Error)
	for _, t := range []model.Tag{coffee, desk} {
		fmt.Printf("tag %v has id %v\n", t.Tag, t.ID)
	}

	event := model.Event{
		UserID:    user.ID,
		Timestamp: time.Now(),
		Notes:     "i am some notes, yo!"}
	fmt.Println(tx.Create(&event).Error)
	fmt.Printf("event %v\n", event.ID)

	tagRefs := make([]model.EventTag, 2)
	tagRefs[0] = model.EventTag{
		EventID: event.ID,
		TagID:   coffee.ID,
		Number:  2}
	tagRefs[1] = model.EventTag{
		EventID: event.ID,
		TagID:   desk.ID,
		Number:  1}
	for _, tr := range tagRefs {
		fmt.Println(tx.Create(&tr).Error)
		fmt.Printf("event tag %v:%v\n", tr.EventID, tr.TagID)
	}

	events := []model.Event{}
	fmt.Println(tx.Preload("Tags").Preload("Tags.Tag").Find(&events).Error)
	for _, e := range events {
		fmt.Printf("event %v: %v %v\n", e.ID, e.Timestamp, e.Tags)
	}
	tags := events[0].Tags
	fmt.Printf("event %v, &tags: %p, &tags[0]: %p, &tags[1]: %p", events[0].ID, &tags, &(tags[0]), &(tags[1]))
	tx.Commit()
}
