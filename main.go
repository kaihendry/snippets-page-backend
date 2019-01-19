package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"snippets.page-backend/config"
	"snippets.page-backend/endpoint"
	"snippets.page-backend/validation"
)

const appConfig = "config.json"

func main() {
	config, err := config.Load(appConfig)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	session, err := mgo.DialWithTimeout(config.Db.ConnectionAddress, time.Duration(15*time.Second))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(config.Db.Database)
	endpoint := endpoint.Endpoint{
		Db:        db,
		JWTSecret: config.JWT.Secret,
	}
	app := echo.New()
	app.Debug = true
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	app.Validator = validation.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//entry point (frontend)
	entryPoint := app.Group("/")
	entryPoint.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "frontend",
		Browse: true,
		HTML5:  true,
	}))
	//API
	public := app.Group("/v1")
	public.POST("/authorization", endpoint.Authorization)
	public.POST("/users", endpoint.CreateUser)
	public.GET("/snippets", endpoint.GetAllPublicSnippets)
	private := app.Group("/v1", middleware.JWT([]byte(endpoint.JWTSecret)))
	private.GET("/me", endpoint.Me)
	private.GET("/me/snippets", endpoint.GetSnippets)
	private.GET("/me/snippets/:id", endpoint.GetSnippet)
	private.GET("/me/snippets/tags", endpoint.GetTags)
	private.POST("/me/snippets", endpoint.CreateSnippet)
	private.PUT("/me/snippets/:id", endpoint.UpdateSnippet)
	private.DELETE("/me/snippets/:id", endpoint.DeleteSnippet)
	if config.TLS.Enable == true {
		app.StartTLS(config.App.Port, config.TLS.Cert, config.TLS.Key)
	} else {
		app.Start(config.App.Port)
	}
}
