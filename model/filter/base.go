package filter

//Base common query filter
type Base struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

//GetLimit returns counts intes per page
func (query *Base) GetLimit() int {
	return query.Limit
}

//GetPage returns page
func (query *Base) GetPage() int {
	return query.Page
}
