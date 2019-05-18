package models

import "gopkg.in/mgo.v2/bson"

type Person struct {
	ID	bson.ObjectId `bson:"_id" json:"id"`
	Firstname string `bson:"firstname" json:"firstname"`
	Lastname string `bson:"lastname" json:"lastname"`
	Address *Address `bson:"address" json:"address"`
}

type Address struct {
	City	string `bson:"city" json:"city"`
	State	string `bson:"state" json:"state"`
}