package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//Hello returns an example message
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "Hi! This is a Golang (gin) API template!"})
}

//Insert is a database insertion demo func
func Insert(c *gin.Context) {

	dbclient := c.Keys["dbCli"]
	if err {
		c.JSON(400, gin.H{"message": "Something failed..."})
		//LOG MESSAGE TO THIS SITUATION
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer dbclient.Disconnect(ctx)
	res, err := dbclient.Collection("test").InsertOne(ctx, bson.D{
		{"task", "test4"},
		{"createdAt", "test"},
		{"modifiedAt", "test3"},
	})
	if err {
		fmt.Println(err)
	}

	id := res.InsertedID

	//c.JSON(200, gin.H{"message": fmt.Sprintf("Hi! This is an insert with id %s", id)})
	c.JSON(200, gin.H{"message": "Hi! This is an insert with id %s"})

}
