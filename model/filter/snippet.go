package filter

import (
	"strings"
)

//Snippet filter for snippets
type Snippet struct {
	Base
	Labels   string `query:"filter[labels]"`
	Language string `query:"filter[language]"`
	Search   string `query:"q"`
	Fields   string `query:"fields"`
}

//NewSnippet init filter for snippets
func NewSnippet() *Snippet {
	filter := &Snippet{}
	filter.Page = 1
	filter.Limit = 50
	filter.Sort = "-created_at"
	return filter
}

//GetLabels returs ,ap
func (s *Snippet) GetLabels() []string {
	if s.Labels == "" {
		return nil
	}
	return strings.Split(s.Labels, ",")
}

//GetLanguage
func (s *Snippet) GetLanguage() string {
	return s.Language
}

//GetKeywords
func (s *Snippet) GetKeywords() string {
	return s.Search
}
