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

//Authorization - endpoint for user auth
func (e *Endpoint) Authorization(context echo.Context) (err error) {
	auth := &model.Authorization{}
	if err = context.Bind(auth); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if err = context.Validate(auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := &model.User{}
	err = e.Db.C("users").Find(bson.M{
		"$or": []bson.M{
			bson.M{"login": auth.LoginOrEmail},
			bson.M{"email": auth.LoginOrEmail},
		},
	}).One(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Access denied")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(auth.Password)); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Access denied")
	}
	//JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = user.Login
	claims["email"] = user.Email
	claims["_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	t, err := token.SignedString([]byte(e.JWTSecret))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Something went wrong...")
	}
	user.Token = t
	return context.JSON(http.StatusOK, user)
}
