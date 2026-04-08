package campaigns

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCampaignsRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewCampaignsRepository(db)
	service := NewCampaignsService(repo)
	handler := NewCampaignsHandler(service)

	campaigns := r.Group("/campaigns")
	{
		campaigns.POST("", handler.CreateCampaigns)
		campaigns.GET("", handler.FindAllCampaigns)
		campaigns.GET("/:id", handler.FindByIDCampaigns)
		campaigns.PUT("/:id", handler.UpdateCampaigns)
		campaigns.DELETE("/:id", handler.DeleteCampaigns)
	}
}
