package model

//Authorization special struct for auth
type Authorization struct {
	LoginOrEmail string `bson:"login_or_email" validate:"required" form:"login_or_email"`
	Password     string `bson:"password" validate:"required"`
}
