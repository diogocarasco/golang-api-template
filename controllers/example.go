package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/diogocarasco/golang-api-template/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//Hello returns an example message
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "Hi! This is a Golang (gin) API template!"})
}

//Insert is a database insertion demo func
func Insert(c *gin.Context) {

	dbclient := c.Keys["dbCli"].(*mongo.Client)
	if dbclient == nil {
		c.JSON(400, gin.H{"message": "Something failed..."})
		//LOG MESSAGE TO THIS SITUATION
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Release resources

	document := models.Examplemodel{
		Field1: "field1",
		Field2: "field2",
	}

	res, err := dbclient.Database("teste").Collection("test").InsertOne(ctx, document)
	if err != nil {
		fmt.Println(err)
	}

	id := res.InsertedID

	c.JSON(200, gin.H{"message": fmt.Sprintf("Hi! This is an insert with id %s", id)})

}

func InsertMany(c *gin.Context) {

	dbclient := c.Keys["dbCli"].(*mongo.Client)
	if dbclient == nil {
		c.JSON(400, gin.H{"message": "Something failed..."})
		//LOG MESSAGE TO THIS SITUATION
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Release resources

	document1 := models.Examplemodel{Field1: "fieldmany1", Field2: "fieldmany2"}
	document2 := models.Examplemodel{Field1: "fieldmany1", Field2: "fieldmany2"}

	documents := []interface{}{document1, document2}

	res, err := dbclient.Database("teste").Collection("test").InsertMany(ctx, documents)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Hi! This is an insert with id %s", res.InsertedIDs...)})

}
