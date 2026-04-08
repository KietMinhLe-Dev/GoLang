package tenants

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTenantsRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewTenantRepository(db)
	service := NewTenantService(repo)
	handler := NewTenantsHandler(service)

	tenants := r.Group("/tenants")
	{
		tenants.POST("", handler.CreateTenant)
		tenants.GET("", handler.FindAllTenants)
		tenants.GET("/:id", handler.FindByIDTenant)
		tenants.PUT("/:id", handler.UpdateTenant)
		tenants.DELETE("/:id", handler.DeleteTenant)
	}
}
