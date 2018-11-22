package endpoint

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"snippets.page-backend/model"
)

//Authentication - endpoint for user auth
func (e *Endpoint) Authentication(context echo.Context) (err error) {
	auth := new(model.Authentication)
	if err = context.Bind(auth); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if err = context.Validate(auth); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	u := new(model.User)
	if err = e.Db.C("users").Find(bson.M{
		"$or": []bson.M{
			bson.M{"login": auth.LoginOrEmail},
			bson.M{"email": auth.LoginOrEmail},
		},
	}).One(&u); err != nil {
		return &echo.HTTPError{Code: http.StatusForbidden, Message: "Something went wrong."}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(auth.Password)); err != nil {
		return &echo.HTTPError{Code: http.StatusForbidden, Message: "Access denied"}
	}
	//JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = u.Login
	claims["email"] = u.Email
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	t, err := token.SignedString([]byte("secret"))
	u.Token = t

	return context.JSON(http.StatusOK, u)
}
