package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type NoteHandler interface {
	CreateNote(ctx *gin.Context)
	FindAllNotes(ctx *gin.Context)
	FindByIDNote(ctx *gin.Context)
	UpdateNote(ctx *gin.Context)
	DeleteNote(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo noteHandler struct {
// 1. service => Nhận về từ file service.go
// 2. noteHandler => Trả về Struct
type noteHandler struct {
	service NoteService
}

// Tạo mới handler
// Cấu trúc khai báo NewNoteHandler(service NoteService) NoteHandler {
// 1. (service NoteService) => Nhận về từ file service.go
// 2. NoteHandler => Trả về Interface
// 3. return &noteHandler{service: service} => Trả về Struct
func NewNoteHandler(service NoteService) NoteHandler {
	return &noteHandler{service: service}
}

// Hàm tạo mới note
func (h *noteHandler) CreateNote(ctx *gin.Context) {
	var req NotesRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leadID, _ := uuid.Parse(req.LeadID)
	userID, _ := uuid.Parse(req.UserID)
	tenantID, _ := uuid.Parse(req.TenantID)

	note := &Note{
		ID:          uuid.New(),
		Content:     &req.Notes,
		LeadID:      &leadID,
		CreatedByID: &userID,
		TenantID:    &tenantID,
	}

	if err := h.service.CreateNote(note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "data": note})
}

// Hàm tìm tất cả notes
func (h *noteHandler) FindAllNotes(ctx *gin.Context) {
	notes, err := h.service.FindAllNotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": notes})
}

// Hàm tìm note theo ID
func (h *noteHandler) FindByIDNote(ctx *gin.Context) {
	id := ctx.Param("id")
	note, err := h.service.FindByIDNote(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": note})
}

// Hàm cập nhật note
func (h *noteHandler) UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req NotesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leadID, _ := uuid.Parse(req.LeadID)
	userID, _ := uuid.Parse(req.UserID)
	tenantID, _ := uuid.Parse(req.TenantID)

	note := &Note{
		ID:          parsedID,
		Content:     &req.Notes,
		LeadID:      &leadID,
		CreatedByID: &userID,
		TenantID:    &tenantID,
	}

	if err := h.service.UpdateNote(note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
}

// Hàm xóa note
func (h *noteHandler) DeleteNote(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.DeleteNote(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
