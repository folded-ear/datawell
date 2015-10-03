/*
Package model houses the structs that are persisted in the database.
*/
package model

import (
	"database/sql"
	"time"
)

/*
User represents a user of datawell which owns events and tags.
*/
type User struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Username  string
	Passhash  string
	Events    []Event
}

/*
Tag represents a single tag in datawell, which are owned by a single user.
*/
type Tag struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
	Tag       string
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
	Tags      []EventTag
	Notes     string
}

/*
EventTag represents a tag that was added to an event, optionally with a number.
*/
type EventTag struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Event     Event
	Tag       Tag
	Number    sql.NullFloat64
}
