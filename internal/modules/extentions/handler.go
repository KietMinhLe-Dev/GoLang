package extentions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type ExtentionHandler interface {
	CreateExtention(ctx *gin.Context)
	FindAllExtentions(ctx *gin.Context)
	FindByIDExtention(ctx *gin.Context)
	UpdateExtention(ctx *gin.Context)
	DeleteExtention(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo extentionHandler struct {
// 1. service => Nhận về từ file service.go
// 2. extentionHandler => Trả về Struct
type extentionHandler struct {
	service ExtentionService
}

// Tạo mới handler
// Cấu trúc khai báo NewExtentionHandler(service ExtentionService) ExtentionHandler {
// 1. (service ExtentionService) => Nhận về từ file service.go
// 2. ExtentionHandler => Trả về Interface
// 3. return &extentionHandler{service: service} => Trả về Struct
func NewExtentionHandler(service ExtentionService) ExtentionHandler {
	return &extentionHandler{service: service}
}

// Hàm tạo mới extention
func (h *extentionHandler) CreateExtention(ctx *gin.Context) {
	var req ExtentionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(req.TenantId)
	userID, _ := uuid.Parse(req.UserId)

	extention := &Extention{
		ID:       uuid.New(),
		Ext:      &req.Ext,
		Password: &req.Password,
		Name:     &req.Name,
		IsRecord: &req.IsRecord,
		TenantID: &tenantID,
		UserID:   &userID,
	}

	if err := h.service.CreateExtention(extention); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Extension created successfully", "data": extention})
}

// Hàm tìm tất cả extentions
func (h *extentionHandler) FindAllExtentions(ctx *gin.Context) {
	extentions, err := h.service.FindAllExtentions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": extentions})
}

// Hàm tìm extention theo ID
func (h *extentionHandler) FindByIDExtention(ctx *gin.Context) {
	id := ctx.Param("id")
	extention, err := h.service.FindByIDExtention(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Extension not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": extention})
}

// Hàm cập nhật extention
func (h *extentionHandler) UpdateExtention(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req ExtentionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(req.TenantId)
	userID, _ := uuid.Parse(req.UserId)

	extention := &Extention{
		ID:       parsedID,
		Ext:      &req.Ext,
		Password: &req.Password,
		Name:     &req.Name,
		IsRecord: &req.IsRecord,
		TenantID: &tenantID,
		UserID:   &userID,
	}

	if err := h.service.UpdateExtention(extention); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Extension updated successfully"})
}

// Hàm xóa extention
func (h *extentionHandler) DeleteExtention(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.DeleteExtention(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Extension deleted successfully"})
}
