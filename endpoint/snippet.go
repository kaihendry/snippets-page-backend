package endpoint

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"snippets.page-backend/model"
	"snippets.page-backend/model/filter"
)

//GetAllPublicSnippets returns all public snippets frrom database
//[GET] /v1/snippets
func (e *Endpoint) GetAllPublicSnippets(context echo.Context) (err error) {
	filter := filter.NewSnippet()
	if err = context.Bind(filter); err != nil {
		return err
	}
	if err = context.Validate(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	conditions := bson.M{}
	conditions["public"] = true
	if tags := filter.GetTags(); tags != nil {
		conditions["tags"] = bson.M{"$in": tags}
	}
	if keywords := filter.GetKeywords(); keywords != "" {
		conditions["$text"] = bson.M{"$search": keywords}
	}
	if favorites := filter.GetFavorites(); favorites == true {
		conditions["favorite"] = true
	}
	pipeline := []bson.M{
		bson.M{"$match": conditions},
		bson.M{"$lookup": bson.M{"from": "users", "localField": "user_id", "foreignField": "_id", "as": "user"}},
		bson.M{"$unwind": "$user"},
		bson.M{"$project": filter.GetFields()},
		bson.M{"$skip": (filter.GetPage() - 1) * filter.GetLimit()},
		bson.M{"$limit": filter.GetLimit()},
		bson.M{"$sort": filter.GetSort()},
	}
	snippets := make([]bson.M, 0)
	total, _ := e.Db.C("snippets").Find(bson.M{"public": true}).Count()
	err = e.Db.C("snippets").Pipe(pipeline).All(&snippets)
	e.addPaginationHeaders(context, filter, total)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}
	return context.JSON(http.StatusOK, &snippets)
}

//GetSnippets returs snippets for current user
//[GET] /v1/me/snippets
func (e *Endpoint) GetSnippets(context echo.Context) (err error) {
	filter := filter.NewSnippet()
	if err = context.Bind(filter); err != nil {
		return err
	}
	if err = context.Validate(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	snippets := make([]bson.M, 0)
	conditions := bson.M{}
	conditions["user_id"] = bson.ObjectIdHex(e.getUserID(context))
	if tags := filter.GetTags(); tags != nil {
		conditions["tags"] = bson.M{"$in": tags}
	}
	if keywords := filter.GetKeywords(); keywords != "" {
		conditions["$text"] = bson.M{"$search": keywords}
	}
	if favorites := filter.GetFavorites(); favorites == true {
		conditions["favorite"] = true
	}
	pipeline := []bson.M{
		bson.M{"$match": conditions},
		bson.M{"$project": filter.GetFields()},
		bson.M{"$skip": (filter.GetPage() - 1) * filter.GetLimit()},
		bson.M{"$limit": filter.GetLimit()},
		bson.M{"$sort": filter.GetSort()},
	}
	err = e.Db.C("snippets").Pipe(pipeline).All(&snippets)
	total, _ := e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context))}).Count()
	e.addPaginationHeaders(context, filter, total)
	return context.JSON(http.StatusOK, snippets)
}

func (e *Endpoint) GetSnippet(context echo.Context) (err error) {
	id := context.Param("id")
	snippet := &model.Snippet{}
	err = e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context)), "_id": bson.ObjectIdHex(id)}).One(&snippet)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return context.JSON(http.StatusOK, snippet)
}

//CreateSnippet creates snippet by current user
//[POST] /v1/me/snippets
func (e *Endpoint) CreateSnippet(context echo.Context) (err error) {
	snippet := &model.Snippet{
		ID:        bson.NewObjectId(),
		UserID:    bson.ObjectIdHex(e.getUserID(context)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = context.Bind(snippet); err != nil {
		return err
	}
	if err = context.Validate(snippet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = e.Db.C("snippets").Insert(snippet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return context.JSON(http.StatusCreated, snippet)
}

//UpdateSnippet updates snippet by id for current user
//[PUT] /me/snippets/<id>
func (e *Endpoint) UpdateSnippet(context echo.Context) (err error) {
	id := context.Param("id")
	snippet := &model.Snippet{}
	err = e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context)), "_id": bson.ObjectIdHex(id)}).One(&snippet)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err = context.Bind(snippet); err != nil {
		return err
	}
	if err = context.Validate(snippet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"favorite":   snippet.Favorite,
			"title":      snippet.Title,
			"files":      snippet.Files,
			"tags":       snippet.Tags,
			"public":     snippet.Public,
			"updated_at": time.Now(),
		}},
		Upsert:    false,
		Remove:    false,
		ReturnNew: true,
	}
	_, err = e.Db.C("snippets").FindId(bson.ObjectIdHex(id)).Apply(change, snippet)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, snippet)
}

//DeleteSnippet removes snippet by _id for current user
//[DELETE] /v1/me/snippets/:id
func (e *Endpoint) DeleteSnippet(context echo.Context) (err error) {
	id := context.Param("id")
	err = e.Db.C("snippets").Remove(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context)), "_id": bson.ObjectIdHex(id)})
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return context.JSON(http.StatusNoContent, nil)
}

//GetTags returns collection of tags
//[PUT] /me/tags/
func (e *Endpoint) GetTags(context echo.Context) (err error) {
	result := make([]string, 0)
	err = e.Db.C("snippets").Find(bson.M{"user_id": bson.ObjectIdHex(e.getUserID(context))}).Distinct("tags", &result)
	if err != nil {
		return context.JSON(http.StatusNoContent, nil)
	}
	return context.JSON(http.StatusOK, result)
}
