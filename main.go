// @title Student Management API
// @version 1.0
// @description A sample API to manage students.
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
	"student-api/config"
	_ "student-api/docs"
	"student-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	routes.StudentRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
