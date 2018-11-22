package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//User model
type User struct {
	ID           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Login        string        `json:"login" bson:"login" validate:"required"`
	PasswordHash string        `json:"-" bson:"password_hash,omitempty"`
	Email        string        `json:"email" bson:"email" validate:"email"`
	Token        string        `json:"token" bson:"token,omitempty"`
	Banned       bool          `json:"banned" bson:"banned"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
}
