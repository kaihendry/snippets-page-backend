package main

import (
	"net/http"

	"github.com/globalsign/mgo"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("v1/authorization", endpoint.Authorization)
	e.POST("v1/users", endpoint.CreateUser)
	e.GET("v1/snippets", endpoint.GetAllPublicSnippets)

	router := e.Group("/v1")
	router.Use(middleware.JWT([]byte("secret")))
	router.GET("/me", endpoint.Me)
	//snipepts
	router.GET("/me/snippets", endpoint.GetSnippets)
	router.POST("/me/snippets", endpoint.CreateSnippet)
	router.PUT("/me/snippets/:id", endpoint.UpdateSnippet)
	router.DELETE("/me/snippets/:id", endpoint.DeleteSnippet)

	e.Start(":8888")
}
