package models

import (
	"time"
)

// InsertTweet
type InsertTweet struct {
	UserId  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
