package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	// Add dependencies like repository or service here
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from example module!",
	})
}
