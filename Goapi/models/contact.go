package models

import "gopkg.in/mgo.v2/bson"

type Contact struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"fistName"bson:"firtsName"`
	LastName  string        `json:"Lastname"bson:"Lastname"`
	Tel       string        `json:"tel"bson:"tel"`
	Email     string        `json:"email"bson:"email"`
}
