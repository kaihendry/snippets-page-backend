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

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	hosts := map[string]*Host{}
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
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	hosts["cloud.snippets.page:8888"] = &Host{app}
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is app")
	})
	api := echo.New()
	api.Debug = true
	api.Validator = validation.New()
	api.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	api.Use(middleware.Logger())
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	hosts["api.snippets.page:8888"] = &Host{api}
	//public
	public := api.Group("/v1")
	public.GET("/ping", func(context echo.Context) error {
		return context.String(http.StatusOK, "pong")
	})
	public.POST("/authorization", endpoint.Authorization)
	public.POST("/users", endpoint.CreateUser)
	public.GET("/snippets", endpoint.GetAllPublicSnippets)
	//private
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
	e.Logger.Fatal(e.Start(":8888"))
}
