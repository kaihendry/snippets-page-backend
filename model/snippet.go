package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Snippet model
type Snippet struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`
	Marked    bool          `json:"marked" bson:"marked"`
	Title     string        `json:"title" bson:"title"`
	Text      string        `json:"text" bson:"text"`
	Language  string        `json:"language" bson:"language"`
	Public    bool          `json:"public" bson:"public"`
	Labels    []string      `json:"labels" bson:"labels"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
