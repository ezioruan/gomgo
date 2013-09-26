package gomgo

import (
	"fmt"
	"labix.org/v2/mgo"
)

var URL = "127.0.0.1"

var session *mgo.Session

func GetSession() *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		fmt.Errorf("getSession error: %v", err)
		return nil
	}
	return session
}

func GetDB() *mgo.Database {
	return GetSession().DB("test")
}
