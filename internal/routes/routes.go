package routes

import (
	"GO/internal/db"
	"GO/internal/modules/appointment"
	calllogs "GO/internal/modules/call_logs"
	"GO/internal/modules/campaigns"
	"GO/internal/modules/extentions"
	"GO/internal/modules/leads"
	"GO/internal/modules/notes"
	"GO/internal/modules/tenants"
	"GO/internal/modules/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// API routes Group
	routes := r.Group("/")

	// Lấy kết nối Database dùng chung
	database := db.DB

	// Đăng ký Router của từng Module
	appointment.SetupAppointmentRouter(routes, database)
	calllogs.SetupCallLogsRouter(routes, database)
	campaigns.SetupCampaignsRouter(routes, database)
	extentions.SetupExtentionsRouter(routes, database)
	leads.SetupLeadsRouter(routes, database)
	notes.SetupNotesRouter(routes, database)
	tenants.SetupTenantsRouter(routes, database)
	users.SetupUsersRouter(routes, database)

	return r
}
