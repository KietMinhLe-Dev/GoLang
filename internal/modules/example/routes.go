package example

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	handler := NewExampleHandler()

	example := r.Group("/example")
	{
		example.GET("/hello", handler.GetHello)
	}
}
