package db

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

type Mongo struct {
	*mgo.Database
}

func InitMongo() *Mongo {
	mongoSession, _ := mgo.Dial("localhost:27017")
	database := mongoSession.DB("snippets_page")
	mdb := &Mongo{database}
	return mdb
}

func (db *Mongo) FindAll(collection string) {

	fmt.Println("Find all from mongo..")
}

func (db *Mongo) FindBy() {

}

func (db *Mongo) FindOneBy() {

}

func (db *Mongo) FindById() {

}

func (db *Mongo) Insert() {

}

func (db *Mongo) Update() {

}

func (db *Mongo) Delete() {

}
