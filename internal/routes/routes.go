package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Dòng này có nghĩa là nếu có lỗi xảy ra thì sẽ tự động recover lại
	r.Use(gin.Recovery())

	return r
}
