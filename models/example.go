package models

type Examplemodel struct {
	Field1 string `json:"jsonfield1" bson:"bsonfield1"`
	Field2 string `json:"jsonfield2" bson:"bsonfield2"`
}
