package filter

import (
	"strings"
)

//Snippet filter for snippets
type Snippet struct {
	Base
	Tags      string `query:"filter[tags]"`
	Search    string `query:"q"`
	Favorites bool   `query:"filter[favorites]"`
	Fields    string `query:"fields"`
	Sort      string `query:"sort" validate:"omitempty,eq=created_at|eq=-created_at"`
}

//NewSnippet init filter for snippets
func NewSnippet() *Snippet {
	filter := &Snippet{}
	filter.Page = 1
	filter.Limit = 100
	filter.Sort = "-created_at"
	filter.Fields = "_id,user_id,favorite,title,files,public,tags,created_at,updated_at"
	return filter
}

func (s *Snippet) GetSort() map[string]int {
	sort := make(map[string]int)
	if string([]rune(s.Sort)[0]) == "-" {
		sort[s.Sort[1:]] = -1
	} else {
		sort["created_at"] = +1
	}
	return sort
}

//GetTags returs
func (s *Snippet) GetTags() []string {
	if s.Tags == "" {
		return nil
	}
	return strings.Split(s.Tags, ",")
}

//GetKeywords
func (s *Snippet) GetKeywords() string {
	return s.Search
}

func (s *Snippet) GetFavorites() bool {
	return s.Favorites
}

//GetFields returns the named elements and attributes associated with resources
func (s *Snippet) GetFields() map[string]int {
	if s.Fields == "" {
		return nil
	}
	query := strings.Split(s.Fields, ",")
	fields := make(map[string]int)
	for _, field := range query {
		fields[field] = 1
	}
	return fields
}
