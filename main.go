package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"snippets.page-backend/endpoint"
	"snippets.page-backend/validation"
)

func main() {
	//config := config.Load("config.json")
	session, err := mgo.DialWithTimeout("mongo:27017", time.Duration(15*time.Second))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("snippets_page")
	endpoint := endpoint.Endpoint{
		Db: db,
	}
	e := echo.New()
	e.Debug = true
	e.Validator = validation.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//public
	router := e.Group("/v1")
	router.GET("/ping", func(context echo.Context) error {
		return context.String(http.StatusOK, "pong")
	})
	router.POST("/authorization", endpoint.Authorization)
	router.POST("/users", endpoint.CreateUser)
	router.GET("/snippets", endpoint.GetAllPublicSnippets)
	//private
	secure := e.Group("/v1", middleware.JWT([]byte("secret")))
	secure.GET("/me", endpoint.Me)
	secure.GET("/me/snippets", endpoint.GetSnippets)
	secure.GET("/me/snippets/:id", endpoint.GetSnippet)
	secure.GET("/me/snippets/tags", endpoint.GetTags)
	secure.POST("/me/snippets", endpoint.CreateSnippet)
	secure.PUT("/me/snippets/:id", endpoint.UpdateSnippet)
	secure.DELETE("/me/snippets/:id", endpoint.DeleteSnippet)

	e.Start(":8888")
}
