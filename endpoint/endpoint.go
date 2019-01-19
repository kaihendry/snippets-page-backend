package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"snippets.page-backend/model/filter"
)

//Endpoint represents HTTP request handler
type Endpoint struct {
	Db        *mgo.Database
	JWTSecret string
}

func (e *Endpoint) getUserID(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["_id"].(string)
}

func (e *Endpoint) addPaginationHeaders(context echo.Context, query filter.Filter, total int) {
	totalPages := 1
	if pages := total / query.GetLimit(); pages > 0 {
		totalPages = pages
	}
	context.Response().Header().Add("X-Pagination-Total-Count", strconv.Itoa(total))
	context.Response().Header().Add("X-Pagination-Page-Count", strconv.Itoa(totalPages))
	context.Response().Header().Add("X-Pagination-Current-Page", strconv.Itoa(query.GetPage()))
	context.Response().Header().Add("X-Pagination-Per-Page", strconv.Itoa(query.GetLimit()))
	lastPage := totalPages
	nextPage := query.GetPage() + 1
	if query.GetPage() == lastPage {
		nextPage = lastPage
	}
	host := context.Request().Host
	scheme := context.Scheme()
	path := context.Request().URL.Path
	self := fmt.Sprintf("<%s://%s%s?page=%d>; rel=self", scheme, host, path, query.GetPage())
	next := fmt.Sprintf("<%s://%s%s?page=%d>; rel=next", scheme, host, path, nextPage)
	last := fmt.Sprintf("<%s://%s%s?page=%d>; rel=last", scheme, host, path, lastPage)
	context.Response().Header().Add("Link", strings.Join([]string{self, next, last}, ","))
}
