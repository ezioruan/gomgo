package gomgo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
)

type BaseModel struct {
	parent         interface{}
	parentType     reflect.Type
	collectionName string
	collection     *mgo.Collection
}

func (self *BaseModel) init() {
	self.parentType = reflect.TypeOf(self.parent)
	self.collection = GetDB().C(self.collectionName)
}

func (self *BaseModel) Save() error {
	v := reflect.ValueOf(self.parent).Elem().Interface()
	fmt.Printf("Save %v", v)
	err := self.collection.Insert(v)
	fmt.Errorf("Save error:%v", err)
	return err
}

type User struct {
	*BaseModel
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Money int           `bson:"money"`
}

func NewUser() *User {
	baseModel := &BaseModel{nil, nil, "user", nil}
	user := &User{}
	baseModel.parent = user
	baseModel.init()
	user.BaseModel = baseModel
	return user
}
