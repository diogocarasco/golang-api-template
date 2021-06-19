package main

import (
	"github.com/diogocarasco/golang-api-template/server"
)

func main() {

	app := server.GetServer()

	app.Run(":8080")

}
