package main

import (
	_ "github.com/maharaj2113/alarm-demo/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()

	// Register the route for Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Your API routes
	router.GET("/ping", pingHandler)

	router.Run(":8080")
}

// @Summary Ping
// @Description Ping handler
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(200, "pong")
}
