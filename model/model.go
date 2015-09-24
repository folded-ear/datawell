package model

import (
	"time"
)

type baseModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	baseModel
	Name     string
	Username string
	Passhash string
	Events   []Event
}

type Tag struct {
	baseModel
	Tag string
}

type Event struct {
	baseModel
	UserID    int64 `sql:"index"`
	Timestamp time.Time
	Tags      []Tag `gorm:"many2many:event_tags;"`
	Notes     string
}
