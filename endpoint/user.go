package endpoint

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"
	"snippets.page-backend/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//CreateUser - registration new user
//[POST] /v1/users
func (e *Endpoint) CreateUser(context echo.Context) (err error) {
	u := &model.User{ID: bson.NewObjectId(), CreatedAt: time.Now()}
	if err = context.Bind(u); err != nil {
		return err
	}
	if err = context.Validate(u); err != nil {
		return err
	}
	user := &model.User{}
	if err = e.Db.C("users").Find(bson.M{"$or": []bson.M{bson.M{"login": u.Login}, bson.M{"email": u.Email}}}).One(user); err == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Email or login already exists"}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(context.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadGateway, Message: "Something went fron..."}
	}
	u.PasswordHash = string(hash)
	err = e.Db.C("users").Insert(u)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadGateway, Message: "Something went fron..."}
	}
	return context.JSON(http.StatusCreated, u)
}

//Me - get the authenticated user
//[GET] /v1/me
func (e *Endpoint) Me(context echo.Context) (err error) {
	authUser := context.Get("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	user := new(model.User)
	if err = e.Db.C("users").FindId(bson.ObjectIdHex(claims["id"].(string))).One(&user); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Something went wrong..."}
	}
	return context.JSON(200, user)
}

//UpdateMe update current user
func (e *Endpoint) UpdateMe(context echo.Context) error {

	q := context.Get("user").(*jwt.Token)
	fmt.Println(q)

	return context.JSON(200, "hello world....")
}
