package main

import (
	"github.com/diogocarasco/golang-api-template/server"
)

// @title Go Gin-Gonic Example API
// @version 1.0
// @description A Gin-Gonic API template for general purposes
// @contact.name Diogo Carasco
// @contact.url http://www.github.com/diogocarasco

func main() {

	app := server.GetServer()

	app.Run(":8080")

}
