package server

import (
	"github.com/diogocarasco/golang-api-template/controllers"
	"github.com/gin-gonic/gin"
)

//SetRoutes Sets the API routes for gin engine
func SetRoutes(server *gin.Engine) {

	server.GET("/", controllers.Hello)
	server.POST("/insert", controllers.Insert)
	server.POST("/insertmany", controllers.InsertMany)

}
