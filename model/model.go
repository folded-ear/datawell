/*
Package model houses the structs that are persisted in the database.
*/
package model

import (
	"time"
)

/*
User represents a user of datawell which owns events and tags.
*/
type User struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username string
	Passhash string
	Name     string

	Events []Event
}

/*
Tag represents a single tag in datawell, which are owned by a single user.
*/
type Tag struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID int64
	Tag    string

	User *User
}

/*
Event represents an event that was recorded in datawell, which are owned by a
single user.
*/
type Event struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID    int64
	Timestamp time.Time
	Notes     string

	Tags []EventTag
	User *User
}

/*
EventTag represents a tag that was added to an event, optionally with a number.
*/
type EventTag struct {
	EventID int64 `gorm:"primary_key"`
	TagID   int64 `gorm:"primary_key"`
	Number  float64

	Event *Event
	Tag   *Tag
}
