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

//CreateUser user registration
//[POST] /v1/users
func (e *Endpoint) CreateUser(context echo.Context) (err error) {
	user := &model.User{ID: bson.NewObjectId(), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err = context.Bind(user); err != nil {
		return err
	}
	if err = context.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	count, _ := e.Db.C("users").Find(bson.M{"$or": []bson.M{bson.M{"login": user.Login}, bson.M{"email": user.Email}}}).Count()
	if count != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Email or login already exists.")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Whoops, something went wrong...")
	}
	user.PasswordHash = string(hash)
	err = e.Db.C("users").Insert(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return context.JSON(http.StatusCreated, user)
}

//Me returns current user
//[GET] /v1/me
func (e *Endpoint) Me(context echo.Context) (err error) {
	gwt := context.Get("user").(*jwt.Token)
	claims := gwt.Claims.(jwt.MapClaims)
	user := &model.User{}
	if err = e.Db.C("users").FindId(bson.ObjectIdHex(claims["id"].(string))).One(&user); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return context.JSON(http.StatusOK, user)
}
