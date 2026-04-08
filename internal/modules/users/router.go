package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUsersRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUsersHandler(service)

	users := r.Group("/users")
	{
		users.POST("", handler.CreateUser)
		users.GET("", handler.FindAllUsers)
		users.GET("/:id", handler.FindByIDUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
