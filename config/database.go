package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/diogocarasco/golang-api-template/utils"
)

// InitializeDB starts the database  connection and appends it to the gin context
func InitializeDB() (*mongo.Client, error) {

	fmt.Print(utils.Blue + "Initializing database -> " + utils.Reset)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connecting to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://golangapi:golangapi@cluster0.92n6j.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	} else {
		//Check the connection
		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println(utils.Red + "Database connection failed!" + utils.Reset)
			cancel()
		}
		fmt.Println(utils.Green + "Database connected!" + utils.Reset)
	}

	return client, err

}
