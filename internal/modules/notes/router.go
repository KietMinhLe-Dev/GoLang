package notes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupNotesRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewNoteRepository(db)
	service := NewNoteService(repo)
	handler := NewNoteHandler(service)

	notes := r.Group("/notes")
	{
		notes.POST("", handler.CreateNote)
		notes.GET("", handler.FindAllNotes)
		notes.GET("/:id", handler.FindByIDNote)
		notes.PUT("/:id", handler.UpdateNote)
		notes.DELETE("/:id", handler.DeleteNote)
	}
}
