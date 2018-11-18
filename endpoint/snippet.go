package endpoint

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"snippets.page-backend/filter"
	"snippets.page-backend/model"

	"github.com/labstack/echo"
)

//FetchSnippets returs snippets for current auth user
func (e *Endpoint) FetchSnippets(context echo.Context) (err error) {
	query := filter.NewQuery()
	if err = context.Bind(query); err != nil {
		return err
	}
	if err = context.Validate(query); err != nil {
		return err
	}
	var snippets []model.Snippet
	e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context))}).Skip((query.Skip - 1) * query.Limit).Limit(query.Limit).All(&snippets)

	totalSnippets, _ := e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context))}).Count()
	totalPages := totalSnippets / query.Limit
	context.Response().Header().Add("X-Pagination-Total-Count", strconv.Itoa(totalSnippets))
	context.Response().Header().Add("X-Pagination-Page-Count", strconv.Itoa(totalPages))
	context.Response().Header().Add("X-Pagination-Current-Page", strconv.Itoa(query.Skip))
	context.Response().Header().Add("X-Pagination-Per-Page", strconv.Itoa(query.Limit))
	//FIX: next
	context.Response().Header().Add("Link", fmt.Sprintf("<http://localhost/users?page=%d>; rel=self, <http://localhost/users?page=%d>; rel=next, <http://localhost/users?page=%d>; rel=last", query.Skip, query.Skip+1, totalPages))

	return context.JSON(200, snippets)
}

//CreateSnippet for current user
func (e *Endpoint) CreateSnippet(context echo.Context) error {

	for i := 0; i < 100; i++ {
		fmt.Println("RAND SNIPPET...")

		snippet := model.Snippet{
			Title:     "Snippet Number" + string(i),
			Text:      "Text" + string(i),
			UserID:    bson.ObjectIdHex("5becbd83f23efbecdd769c40"),
			Language:  "JS",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := e.Db.C("snippets").Insert(snippet)

		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println("end..")

	return context.JSON(200, "dwelrwel;rwe")
}

func (e *Endpoint) UpdateSnippet(context echo.Context) error {
	return context.JSON(200, "test")
}
