package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"snippets.page-backend/db"
	"snippets.page-backend/handler"
)

func main() {
	//config := config.Load("config.json")
	mongoDb := db.InitMongo()
	handler := handler.Handler{
		Db: mongoDb,
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// v1/user/login

	// v1/users
	// v1/snippets
	// v1/tags

	//open actions
	e.POST("v1/authentication", handler.Authentication)

	e.POST("v1/users", func(context echo.Context) error {
		return context.String(200, "this is registration...")
	})

	router := e.Group("/v1")
	router.Use(middleware.JWT([]byte("secret")))
	//Get metadata about the current user.
	router.GET("/me", func(context echo.Context) error {
		return context.String(200, "this is only for auth....")
	})
	//Get snippets for current user.
	router.GET("/me/snippets", func(context echo.Context) error {
		return context.String(200, "this is snippets.")
	})

	fmt.Println(handler)

	e.Start(":8888")
}

func getJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Logined user"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, err
}
