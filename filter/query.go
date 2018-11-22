package filter

//QueryParam returs filtered and validated GET params
type QueryParam struct {
	Sort  string `json:"sort" bson:"sort,omitempty"`
	Limit int    `json:"limit" bson:"limit,omitempty" validate:"gte=1,lte=100"`
	Page  int    `json:"skip" bson:"skip,omitempty"`
}

func NewQuery() *QueryParam {
	query := new(QueryParam)
	query.Limit = 50
	query.Page = 1
	return query
}
