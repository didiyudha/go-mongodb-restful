package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

// GetMongoSession return pointer to mongodb session
func GetMongoSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Println("Error connection to mongodb: ", err.Error())
		panic(err.Error())
	}
	return session
}
