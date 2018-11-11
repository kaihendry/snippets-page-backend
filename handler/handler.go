package handler

import "snippets.page-backend/db"

type Handler struct {
	Db *db.Mongo
}
