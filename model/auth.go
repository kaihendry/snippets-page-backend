package model

//Authentication - special
type Authentication struct {
	LoginOrEmail string `bson:"login_or_email" validate:"required" form:"login_or_email"`
	Password     string `bson:"password" validate:"required"`
}
