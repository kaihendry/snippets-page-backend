package endpoint

import (
	"fmt"
	"strconv"

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
	totalPages := total / query.GetLimit()
	context.Response().Header().Add("X-Pagination-Total-Count", strconv.Itoa(total))
	context.Response().Header().Add("X-Pagination-Page-Count", strconv.Itoa(totalPages))
	context.Response().Header().Add("X-Pagination-Current-Page", strconv.Itoa(query.GetPage()))
	context.Response().Header().Add("X-Pagination-Per-Page", strconv.Itoa(query.GetLimit()))
	lastPage := totalPages
	nextPage := query.GetPage() + 1
	if query.GetPage() == lastPage {
		nextPage = lastPage
	}
	link := fmt.Sprintf("<http://localhost/users?page=%d>; rel=self, <http://localhost/users?page=%d>; rel=next, <http://localhost/users?page=%d>; rel=last", query.GetPage(), nextPage, lastPage)
	context.Response().Header().Add("Link", link)

}
