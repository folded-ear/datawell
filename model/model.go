/*
Package model houses the structs that are persisted in the database.
*/
package model

import (
	"time"
)

/*
baseModel is a struct that declares the core meta fields that all persistent
structs will have via Go's implicit key delegation.
*/
type baseModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
User represents a user of the system.
*/
type User struct {
	baseModel
	Name     string
	Username string
	Passhash string
	Events   []Event
}

/*
Tag represents a single tag in the system.  Note that tags are top-level items;
they are not owned by a specific user.
*/
type Tag struct {
	baseModel
	Tag string
}

/*
Event represents an event that was recorded in datawell.
*/
type Event struct {
	baseModel
	UserID    int64
	Timestamp time.Time
	Tags      []Tag `gorm:"many2many:event_tags;"`
	Notes     string
}
