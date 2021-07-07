package server

import (
	"github.com/diogocarasco/golang-api-template/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

//SetRoutes Sets the API routes for gin engine
func SetRoutes(server *gin.Engine) {

	server.GET("/", controllers.Hello)
	server.POST("/insert", controllers.Insert)
	server.POST("/insertmany", controllers.InsertMany)
	server.GET("/find/:id", controllers.Find)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}
