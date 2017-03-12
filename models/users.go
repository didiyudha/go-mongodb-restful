package models

import (
	"github.com/didiyudha/go-mongodb-restful/db"
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
}

// NewUser return pinter to User struc
func NewUser() *User {
	return &User{}
}

// Save user document
func (u *User) Save() error {
	session := db.GetMongoSession()
	defer session.Close()
	u.ID = bson.NewObjectId()
	return session.DB("blog").C("users").Insert(u)
}

// FindAll get all users from mongodb
func (u *User) FindAll() ([]User, error) {
	session := db.GetMongoSession()
	defer session.Close()
	var users []User
	err := session.DB("blog").C("users").Find(bson.M{}).All(&users)
	return users, err
}

// FindByID find a user from mongodb
func (u *User) FindByID(ID bson.ObjectId) (User, error) {
	sess := db.GetMongoSession()
	defer sess.Close()
	var usr = User{}
	err := sess.DB("blog").C("users").FindId(ID).One(&usr)
	return usr, err
}

// Update a user from mongodb
func (u *User) Update(ID bson.ObjectId) error {
	sess := db.GetMongoSession()
	defer sess.Close()
	return sess.DB("blog").C("users").Update(bson.M{"_id": ID}, u)
}

// Delete a user from mongodb
func (u *User) Delete(ID bson.ObjectId) error {
	sess := db.GetMongoSession()
	defer sess.Close()
	return sess.DB("blog").C("users").Remove(bson.M{"_id": ID})
}
