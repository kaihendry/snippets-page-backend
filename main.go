package main

import (
	"gopkg.in/mgo.v2"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"snippets.page-backend/endpoint"
	"snippets.page-backend/validation"
)

func main() {
	//config := config.Load("config.json")
	session, _ := mgo.Dial("localhost:27017")

	db := session.DB("snippets_page")

	endpoint := endpoint.Endpoint{
		Db: db,
	}
	e := echo.New()
	e.Debug = true
	e.Validator = validation.New()

	e.POST("v1/authentication", endpoint.Authentication)
	e.POST("v1/users", endpoint.CreateUser)

	router := e.Group("/v1")
	router.Use(middleware.JWT([]byte("secret")))
	//Get metadata about the current user.
	router.GET("/me", endpoint.CurrentUser)
	//Get snippets for current user.
	router.GET("/me/snippets", func(context echo.Context) error {
		return context.String(200, "this is snippets.")
	})

	e.Start(":8888")
}
