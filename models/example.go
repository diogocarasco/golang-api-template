package models

type examplemodel struct {
	field1 string `json:"jsonfield1" bson:"bsonfield1"`
	field2 string `json:"jsonfield2" bson:"bsonfield2"`
}
