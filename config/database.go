package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitializeDB starts the database  connection and appends it to the gin context
func InitializeDB() (*mongo.Client, error) {

	fmt.Print("Initializing database -> ")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connecting to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"",
	))
	if err != nil {
		log.Fatal(err)
	} else {
		//Check the connection
		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			cancel()
		}
		fmt.Println("Database connected!")
	}
	//var db type *utils.Database{client,ctx}
	//db := utils.Database{dbclient: client, dbcontext: ctx}

	return client, err

}
