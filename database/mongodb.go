package database

import mgo "gopkg.in/mgo.v2"

// MgoSession is the Global Mongo session
var MgoSession *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	MgoSession = session
}
