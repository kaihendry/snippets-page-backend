package endpoint

import mgo "gopkg.in/mgo.v2"

type Endpoint struct {
	Db *mgo.Database
}
