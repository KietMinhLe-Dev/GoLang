package calllogs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCallLogsRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewCallLogsRepository(db)
	service := NewCallLogsService(repo)
	handler := NewCallLogsHandler(service)

	callLogs := r.Group("/call-logs")
	{
		callLogs.POST("", handler.CreateCallLogs)
		callLogs.GET("", handler.FindAllCallLogs)
		callLogs.GET("/:id", handler.FindByIDCallLogs)
		callLogs.PUT("/:id", handler.UpdateCallLogs)
		callLogs.DELETE("/:id", handler.DeleteCallLogs)
	}
}
