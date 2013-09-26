package gomgo

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

//func TestSession(t *testing.T) {
//	session = GetSession()
//	session.DB("test").C("test").Insert(bson.M{"name": "ezio"})
//}

func TestSave(t *testing.T) {
	user := NewUser()
	user.Name = "ezio"
	user.Money = 10
	user.Id = bson.NewObjectId()
	user.Save()
}
