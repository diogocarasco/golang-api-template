package middleware

import (
	"github.com/diogocarasco/golang-api-template/config"
	"github.com/gin-gonic/gin"
)

func InitializeDB() gin.HandlerFunc {

	client, _ := config.InitializeDB()

	return func(c *gin.Context) { // Appending database client to gin context
		c.Set("dbCli", client)
		c.Next()
	}

}
