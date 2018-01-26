package models

import "gopkg.in/mgo.v2/bson"

type Vegetable struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `bson:"title" json:"title"`
	Select string        `bson:"select" json:"select"`
	Store  string        `bson:"store" json:"store"`
	Ripen  string        `bson:"ripen" json:"ripen"`
	Season string        `bson:"season" json:"season"`
	Image  string        `bson:"image" json:"image"`
	Count  int           `bson:"count" json:"count"`
}
