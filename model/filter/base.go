package filter

import "strings"

//Base common query filter
type Base struct {
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Sort  string `query:"sort"`
}

//GetLimit returns counts intes per page
func (query *Base) GetLimit() int {
	return query.Limit
}

//GetPage returns page
func (query *Base) GetPage() int {
	return query.Page
}

//GetSort returns map with fields
func (query *Base) GetSort() map[string]int {
	if query.Sort == "" {
		return nil
	}
	sort := make(map[string]int)
	fields := strings.Split(query.Sort, ",")
	for _, field := range fields {
		if string([]rune(field)[0]) == "-" {
			sort[field[1:]] = -1
		} else if string([]rune(field)[0]) == "+" {
			sort[field[1:]] = +1
		} else {
			sort[field] = +1
		}
	}
	return sort
}
