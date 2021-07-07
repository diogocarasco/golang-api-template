package models

type Examplemodel struct {
	ExampleId     string `json:"example_id" bson:"example_id" binding:"required" example:"9999"`
	ExampleString string `json:"example_string" bson:"example_string" example:"Example string"`
	ExampleNumber int    `json:"example_number" bson:"example_number" example:"Example number"`
}

type ExampleManyModel struct {
	Examples []Examplemodel `json:"examples" `
}
