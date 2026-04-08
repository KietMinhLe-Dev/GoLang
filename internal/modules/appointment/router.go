package appointment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAppointmentRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewAppointmentRepository(db)
	service := NewAppointmentService(repo)
	handler := NewAppointmentHandler(service)

	appointments := r.Group("/appointments")
	{
		appointments.POST("", handler.CreateAppointment)
		appointments.GET("", handler.FindAllAppointments)
		appointments.GET("/:id", handler.FindByIDAppointment)
		appointments.PUT("/:id", handler.UpdateAppointment)
		appointments.DELETE("/:id", handler.DeleteAppointment)
	}
}
