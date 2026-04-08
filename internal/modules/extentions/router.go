package extentions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupExtentionsRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewExtentionRepository(db)
	service := NewExtentionService(repo)
	handler := NewExtentionHandler(service)

	extentions := r.Group("/extensions")
	{
		extentions.POST("", handler.CreateExtention)
		extentions.GET("", handler.FindAllExtentions)
		extentions.GET("/:id", handler.FindByIDExtention)
		extentions.PUT("/:id", handler.UpdateExtention)
		extentions.DELETE("/:id", handler.DeleteExtention)
	}
}
