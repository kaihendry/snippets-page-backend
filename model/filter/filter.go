package filter

//Filter interface
type Filter interface {
	GetLimit() int
	GetPage() int
}
