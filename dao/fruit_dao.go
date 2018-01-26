package dao

import (
	"log"

	. "just-ripe/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FruitsDAO struct {
	Server   string
	Database string
	DialInfo *mgo.DialInfo
}

var db *mgo.Database

const (
	FRUITCOLLECTION     = "Fruit"
	VEGETABLECOLLECTION = "Vegetable"
	NEWCOLLECTION       = "Fruits"
)

func (m *FruitsDAO) Connect() {
	session, err := mgo.DialWithInfo(m.DialInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
func (m *FruitsDAO) FindAllFruit() ([]Fruit, error) {
	var fruits []Fruit
	err := db.C(FRUITCOLLECTION).Find(bson.M{}).All(&fruits)
	return fruits, err
}

func (m *FruitsDAO) FindAllVegetables() ([]Vegetable, error) {
	var veggies []Vegetable
	err := db.C(VEGETABLECOLLECTION).Find(bson.M{}).All(&veggies)
	return veggies, err
}

func (m *FruitsDAO) FindFruitById(id string) (Fruit, error) {
	var fruit Fruit
	err := db.C(FRUITCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&fruit)
	return fruit, err
}

func (m *FruitsDAO) FindVegetableById(id string) (Vegetable, error) {
	var veggie Vegetable
	err := db.C(VEGETABLECOLLECTION).FindId(bson.ObjectIdHex(id)).One(&veggie)
	return veggie, err
}

func (m *FruitsDAO) UpdateFruit(fruit Fruit) error {
	err := db.C(FRUITCOLLECTION).UpdateId(fruit.ID, &fruit)
	return err
}

func (m *FruitsDAO) UpdateVegetable(veggie Vegetable) error {
	err := db.C(VEGETABLECOLLECTION).UpdateId(veggie.ID, &veggie)
	return err
}
