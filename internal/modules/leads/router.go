package leads

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLeadsRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewLeadRepository(db)
	service := NewLeadService(repo)
	handler := NewLeadsHandler(service)

	leads := r.Group("/leads")
	{
		leads.POST("", handler.CreateLead)
		leads.GET("", handler.FindAllLeads)
		leads.GET("/:id", handler.FindByIDLead)
		leads.PUT("/:id", handler.UpdateLead)
		leads.DELETE("/:id", handler.DeleteLead)
	}
}
