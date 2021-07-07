package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/diogocarasco/golang-api-template/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	document := models.Examplemodel{}
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := dbclient.Database("teste").Collection("test").InsertOne(ctx, document)
	if err != nil {
		fmt.Println(err)
	}

	id := res.InsertedID

	c.JSON(http.StatusOK, gin.H{"message": id})

}

//InsertMany is a database multiple insertion demo func
func InsertMany(c *gin.Context) {

	dbclient := c.Keys["dbCli"].(*mongo.Client)
	if dbclient == nil {
		c.JSON(400, gin.H{"message": "Something failed..."})
		//LOG MESSAGE TO THIS SITUATION
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Release resources

	// document1 := models.Examplemodel{ExampleId: "1", ExampleString: "Example string 1", ExampleNumber: 11}
	// document2 := models.Examplemodel{ExampleId: "2", ExampleString: "Example string 2", ExampleNumber: 22}

	documentList := models.ExampleManyModel{}
	if err := c.ShouldBindJSON(&documentList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	documents := []interface{}{}
	for _, doc := range documentList.Examples {

		documents = append(documents, doc)
	}

	res, err := dbclient.Database("teste").Collection("test").InsertMany(ctx, documents)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message": res.InsertedIDs})

}

//Find is a database read demo func
func Find(c *gin.Context) {

	var results []*models.Examplemodel

	dbclient := c.Keys["dbCli"].(*mongo.Client)
	if dbclient == nil {
		c.JSON(400, gin.H{"message": "Something failed..."})
		//LOG MESSAGE TO THIS SITUATION
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Release resources

	id := c.Param("id")
	filter := bson.D{{Key: "example_id", Value: id}}

	cur, err := dbclient.Database("teste").Collection("test").Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {

		var result models.Examplemodel
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &result)

	}

	c.JSON(200, gin.H{"message": results})

}
