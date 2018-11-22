package endpoint

import (
	"fmt"
	"strconv"

	"snippets.page-backend/filter"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
)

//Endpoint represents HTTP request handler
type Endpoint struct {
	Db *mgo.Database
}

func (e *Endpoint) getUserID(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func (e *Endpoint) addPaginationHeaders(context echo.Context, query *filter.QueryParam, total int) {
	totalPages := total / query.Limit
	context.Response().Header().Add("X-Pagination-Total-Count", strconv.Itoa(total))
	context.Response().Header().Add("X-Pagination-Page-Count", strconv.Itoa(totalPages))
	context.Response().Header().Add("X-Pagination-Current-Page", strconv.Itoa(query.Page))
	context.Response().Header().Add("X-Pagination-Per-Page", strconv.Itoa(query.Limit))
	lastPage := totalPages
	nextPage := query.Page + 1
	if query.Page == lastPage {
		nextPage = lastPage
	}
	link := fmt.Sprintf("<http://localhost/users?page=%d>; rel=self, <http://localhost/users?page=%d>; rel=next, <http://localhost/users?page=%d>; rel=last", query.Page, nextPage, lastPage)
	context.Response().Header().Add("Link", link)

}
