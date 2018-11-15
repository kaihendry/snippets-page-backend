package model

import (
	"gopkg.in/mgo.v2/bson"
)

//Snippet model
type Snippet struct {
	ID       bson.ObjectId
	UserID   bson.ObjectId
	Title    string
	Text     string
	Language string
	Public   bool
	Tags     []string
}
