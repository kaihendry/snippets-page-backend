package model

//Authorization special struct for auth
type Authorization struct {
	LoginOrEmail string `json:"login" bson:"login" validate:"required" form:"login"`
	Password     string `json:"password" bson:"password" validate:"required" form:"password"`
}
