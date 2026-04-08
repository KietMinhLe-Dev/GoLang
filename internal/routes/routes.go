package routes

import (
	"net/http"

	"GO/internal/modules/example"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Middleware
	r.Use(gin.Recovery())

	// Example routes
	example.RegisterRoutes(r)

	// Basic route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
