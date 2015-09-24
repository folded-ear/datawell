package model

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Username  string
	Passhash  string
	Events    []Event
}

type Tag struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tag       string
}

type Event struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64 `sql:"index"`
	Timestamp time.Time
	Tags      []Tag `gorm:"many2many:event_tags;"`
	Notes     string
}
