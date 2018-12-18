package filter

import (
	"strings"
)

//Snippet filter for snippets
type Snippet struct {
	Base
	Tags      string `query:"filter[tags]"`
	Search    string `query:"q"`
	Favorites bool   `query:"favorites"`
	Fields    string `query:"fields"`
}

//NewSnippet init filter for snippets
func NewSnippet() *Snippet {
	filter := &Snippet{}
	filter.Page = 1
	filter.Limit = 50
	filter.Sort = "-created_at"
	return filter
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
