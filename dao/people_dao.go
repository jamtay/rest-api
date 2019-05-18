package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	. "rest-api/models")

type PeopleDAO struct {
	Server string
	Database string
}

var peopleDb *mgo.Database

const (
	PEOPLE_COLLECTION = "people"
)

// https://tour.golang.org/methods/1 explains what (m *PeopleDAO) is
func (m *PeopleDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	peopleDb = session.DB(m.Database)
}


func (p *PeopleDAO) FindAllPeople() ([]Person, error) {
	var people []Person
	err := peopleDb.C(PEOPLE_COLLECTION).Find(bson.M{}).All(&people)
	return people, err
}

func (p *PeopleDAO) FindById(id string) (Person, error) {
	var person = Person{}
	err := peopleDb.C(PEOPLE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&person)
	return person, err
}

func (p *PeopleDAO) Insert(person Person) error {
	err := peopleDb.C(PEOPLE_COLLECTION).Insert(&person)
	return err
}

func (m *PeopleDAO) Delete(person Person) error {
	err := db.C(PEOPLE_COLLECTION).Remove(&person)
	return err
}
