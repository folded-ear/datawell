/*
Package model houses the structs that are persisted in the database.
*/
package model

import (
	"database/sql"
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
User represents a user of datawell which owns events and tags.
*/
type User struct {
	baseModel
	Name     string
	Username string
	Passhash string
	Events   []Event
}

/*
Tag represents a single tag in datawell, which are owned by a single user.
*/
type Tag struct {
	baseModel
	UserID int64
	Tag    string
}

/*
Event represents an event that was recorded in datawell, which are owned by a
single user.
*/
type Event struct {
	baseModel
	UserID    int64
	Timestamp time.Time
	Tags      []EventTag
	Notes     string
}

/*
tagRef is a struct for references of a tag from "something" else, which will be
expressed as embedding structs.
*/
type tagRef struct {
	baseModel
	Tag    Tag
	Number sql.NullFloat64
}

/*
EventTag represents a tag that was added to an event, optionally with a number.
*/
type EventTag struct {
	tagRef
	Event Event
}
