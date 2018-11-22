package endpoint

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
)

type Endpoint struct {
	Db *mgo.Database
}

func (e *Endpoint) getUserID(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
