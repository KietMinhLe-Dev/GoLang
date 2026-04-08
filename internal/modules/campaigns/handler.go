package campaigns

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type CampaignsHandler interface {
	CreateCampaigns(ctx *gin.Context)
	FindAllCampaigns(ctx *gin.Context)
	FindByIDCampaigns(ctx *gin.Context)
	UpdateCampaigns(ctx *gin.Context)
	DeleteCampaigns(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo campaignsHandler struct {
// 1. service => Nhận về từ file service.go
// 2. campaignsHandler => Trả về Struct
type campaignsHandler struct {
	service CampaignsService
}

// Tạo mới handler
// Cấu trúc khai báo NewCampaignsHandler(service CampaignsService) CampaignsHandler {
// 1. (service CampaignsService) => Nhận về từ file service.go
// 2. CampaignsHandler => Trả về Interface
// 3. return &campaignsHandler{service: service} => Trả về Struct
func NewCampaignsHandler(service CampaignsService) CampaignsHandler {
	return &campaignsHandler{service: service}
}

// Hàm tạo mới campaign
// Cấu trúc khai báo func (h *campaignsHandler) CreateCampaigns(ctx *gin.Context) {
// 1. (h *campaignsHandler) => Nhận về từ file handler.go
// 2. http.StatusCreated => Trả về mã lỗi 201
// 3. gin.H{} => Trả về JSON
func (h *campaignsHandler) CreateCampaigns(ctx *gin.Context) {
	var req CampaignRequest // Khai báo biến req kiểu CampaignRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi từ DTO sang Model
	var tenantID *uuid.UUID
	if req.TenantID != nil {
		id, _ := uuid.Parse(*req.TenantID)
		tenantID = &id
	}

	campaign := &Campaign{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
		TenantID:    tenantID,
	}

	// Tạo mới campaign bằng service
	if err := h.service.CreateCampaigns(campaign); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusCreated, gin.H{"message": "Campaign created successfully", "data": campaign})
}

// Hàm tìm tất cả campaigns
// Cấu trúc khai báo func (h *campaignsHandler) FindAllCampaigns(ctx *gin.Context) {
// 1. (h *campaignsHandler) => Nhận về từ file handler.go
// 2. http.StatusOK => Trả về mã lỗi 200
// 3. gin.H{} => Trả về JSON
func (h *campaignsHandler) FindAllCampaigns(ctx *gin.Context) {
	// Tìm tất cả campaigns
	campaigns, err := h.service.FindAllCampaigns()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": campaigns})
}

// Hàm tìm campaign theo ID
// Cấu trúc khai báo func (h *campaignsHandler) FindByIDCampaigns(ctx *gin.Context) {
// 1. (h *campaignsHandler) => Nhận về từ file handler.go
// 2. ctx.Param("id") => Nhận ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *campaignsHandler) FindByIDCampaigns(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Tìm campaign theo ID
	campaign, err := h.service.FindByIDCampaigns(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": campaign})
}

// Hàm cập nhật campaign
// Cấu trúc khai báo func (h *campaignsHandler) UpdateCampaigns(ctx *gin.Context) {
// 1. (h *campaignsHandler) => Nhận về từ file handler.go
// 2. uuid.Parse(id) => Parse ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *campaignsHandler) UpdateCampaigns(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req CampaignRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi từ DTO sang Model
	var tenantID *uuid.UUID
	if req.TenantID != nil {
		id, _ := uuid.Parse(*req.TenantID)
		tenantID = &id
	}

	campaign := &Campaign{
		ID:          parsedID,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
		TenantID:    tenantID,
	}

	// Cập nhật campaign bằng service
	if err := h.service.UpdateCampaigns(campaign); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Campaign updated successfully"})
}

// Hàm xóa campaign
// Cấu trúc khai báo func (h *campaignsHandler) DeleteCampaigns(ctx *gin.Context) {
// 1. (h *campaignsHandler) => Nhận về từ file handler.go
// 2. h.service.DeleteCampaigns(id) => Gọi hàm xóa từ service
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *campaignsHandler) DeleteCampaigns(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Xóa campaign bằng service
	if err := h.service.DeleteCampaigns(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
