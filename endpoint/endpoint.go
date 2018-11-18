package endpoint

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

type Endpoint struct {
	Db *mgo.Database
}

func (e *Endpoint) getUserID(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
