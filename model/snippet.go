package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Snippet represents resources in JSON format.
type Snippet struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   bson.ObjectId `json:"user_id" bson:"user_id"`
	Favorite bool          `json:"favorite" bson:"favorite"`
	Title    string        `json:"title" bson:"title" validate:"required,gte=1,lte=40"`
	Files    []struct {
		Name      string      `json:"name" bson:"name"`
		Content   string      `json:"content" bson:"content"`
		Options   interface{} `json:"options" bson:"options"`
		CreatedAt time.Time   `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
	} `json:"files" bson:"files"`
	Public    bool      `json:"public" bson:"public"`
	Tags      []string  `json:"tags" bson:"tags"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
