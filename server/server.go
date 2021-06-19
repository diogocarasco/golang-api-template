package server

import (
	"github.com/diogocarasco/golang-api-template/middleware"
	"github.com/diogocarasco/golang-api-template/utils"
	"github.com/gin-gonic/gin"
)

// GetServer returns the default gin engine
func GetServer() *gin.Engine {

	server := utils.Startup()
	server.Use(middleware.InitializeDB())
	SetRoutes(server)

	return server

}
