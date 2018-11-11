package model

type base interface {
	FindAll(collection string)
	FindBy()
	FindOneBy()
	FindById()
	Insert()
	Update()
	Delete()
}
