package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//User represents resources in JSON format.
type User struct {
	ID           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Login        string        `json:"login" bson:"login" validate:"required,gte=3,lte=40"`
	PasswordHash string        `json:"-" bson:"password_hash,omitempty"`
	Email        string        `json:"email" bson:"email" validate:"required,email"`
	Token        string        `json:"token" bson:"-"`
	Banned       bool          `json:"banned" bson:"banned"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
	//only for registration
	Password string `json:"-" bson:"-" validate:"required,gte=6,lte=50"`
}
