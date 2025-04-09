package server

import (
	"github.com/adrien3d/fizzbuzz/handlers"
	"github.com/gin-gonic/gin"
)


func SetupRouter() error {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	api := router.Group("/api")
	{
		// Inject dependencies into the handlers
		api.GET("/fizzbuzz", handlers.FizzBuzzHandler)
		api.GET("/stats", handlers.StatsHandler)
	}

	return router.Run(":8080")
}