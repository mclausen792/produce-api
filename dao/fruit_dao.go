package dao

import (
	"log"

	. "github.com/mclausen792/produce-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FruitsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "fruit"
)

func (m *FruitsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
func (m *FruitsDAO) FindAll() ([]Fruit, error) {
	var fruits []Fruit
	err := db.C(COLLECTION).Find(bson.M{}).All(&fruits)
	return fruits, err
}

func (m *FruitsDAO) FindById(id string) (Fruit, error) {
	var fruit Fruit
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&fruit)
	return fruit, err
}

func (m *FruitsDAO) Update(fruit Fruit) error {
	err := db.C(COLLECTION).UpdateId(fruit.ID, &fruit)
	return err
}
