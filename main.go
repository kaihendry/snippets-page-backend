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
	e.GET("v1/snippets", func(context echo.Context) error {
		return context.JSON(200, "all public snippets")
	})
	e.GET("v1/users", func(context echo.Context) error {
		return context.JSON(200, "All users....")
	})

	e.GET("v1/users/:userId/snippets", func(context echo.Context) error {
		return context.JSON(200, "get all snippets by user")
	})

	router := e.Group("/v1")
	router.Use(middleware.JWT([]byte("secret")))
	router.GET("/me", endpoint.Me)
	router.PUT("/me", endpoint.UpdateMe)
	//snipepts
	router.GET("/me/snippets", endpoint.FetchSnippets)
	router.POST("/me/snippets", endpoint.CreateSnippet)
	router.PUT("/me/snippets/", endpoint.UpdateSnippet)

	e.Start(":8888")
}
