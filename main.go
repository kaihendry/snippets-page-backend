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

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	hosts := map[string]*Host{}
	config, err := config.Load("config.json")
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
		Db: db,
	}
	app := echo.New()
	hosts[config.App.Frontend] = &Host{app}
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is app")
	})
	api := echo.New()
	api.Validator = validation.New()
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	hosts[config.App.Backend] = &Host{api}
	api.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Welcome to RESTful API service")
	})
	public := api.Group("/v1")
	public.POST("/authorization", endpoint.Authorization)
	public.POST("/users", endpoint.CreateUser)
	public.GET("/snippets", endpoint.GetAllPublicSnippets)
	private := api.Group("/v1", middleware.JWT([]byte("secret")))
	private.GET("/me", endpoint.Me)
	private.GET("/me/snippets", endpoint.GetSnippets)
	private.GET("/me/snippets/:id", endpoint.GetSnippet)
	private.GET("/me/snippets/tags", endpoint.GetTags)
	private.POST("/me/snippets", endpoint.CreateSnippet)
	private.PUT("/me/snippets/:id", endpoint.UpdateSnippet)
	private.DELETE("/me/snippets/:id", endpoint.DeleteSnippet)
	//server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	e.Debug = true
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]
		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}
		return
	})
	if config.TLS.Enable == true {
		e.StartTLS(config.App.Port, config.TLS.Cert, config.TLS.Key)
	} else {
		e.Start(config.App.Port)
	}
}
