package leads

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type LeadsHandler interface {
	CreateLead(ctx *gin.Context)
	FindAllLeads(ctx *gin.Context)
	FindByIDLead(ctx *gin.Context)
	UpdateLead(ctx *gin.Context)
	DeleteLead(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo leadsHandler struct {
// 1. service => Nhận về từ file service.go
// 2. leadsHandler => Trả về Struct
type leadsHandler struct {
	service LeadService
}

// Tạo mới handler
// Cấu trúc khai báo NewLeadsHandler(service LeadService) LeadsHandler {
// 1. (service LeadService) => Nhận về từ file service.go
// 2. LeadsHandler => Trả về Interface
// 3. return &leadsHandler{service: service} => Trả về Struct
func NewLeadsHandler(service LeadService) LeadsHandler {
	return &leadsHandler{service: service}
}

// Hàm tạo mới lead
// Cấu trúc khai báo func (h *leadsHandler) CreateLead(ctx *gin.Context) {
// 1. (h *leadsHandler) => Nhận về từ file handler.go
// 2. http.StatusCreated => Trả về mã lỗi 201
// 3. gin.H{} => Trả về JSON
func (h *leadsHandler) CreateLead(ctx *gin.Context) {
	var req LeadsRequest // Khai báo biến req kiểu LeadsRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lead := &Lead{
		ID:         uuid.New(),
		Name:       &req.Name,
		Phone:      &req.Phone,
		Address:    &req.Address,
		Source:     &req.Source,
		Stage:      &req.Stage,
		Notes:      &req.Notes,
		UserID:     req.UserID,
		AssignedAt: req.AssignedAt,
		CampaignID: req.CampaignID,
		TenantID:   req.TenantID,
		ShowAt:     req.ShowAt,
		LastCallAt: req.LastCallAt,
	}

	// Tạo mới lead bằng service
	if err := h.service.CreateLead(lead); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusCreated, gin.H{"message": "Lead created successfully", "data": lead})
}

// Hàm tìm tất cả leads
// Cấu trúc khai báo func (h *leadsHandler) FindAllLeads(ctx *gin.Context) {
// 1. (h *leadsHandler) => Nhận về từ file handler.go
// 2. http.StatusOK => Trả về mã lỗi 200
// 3. gin.H{} => Trả về JSON
func (h *leadsHandler) FindAllLeads(ctx *gin.Context) {
	// Tìm tất cả leads
	leads, err := h.service.FindAllLeads()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": leads})
}

// Hàm tìm lead theo ID
// Cấu trúc khai báo func (h *leadsHandler) FindByIDLead(ctx *gin.Context) {
// 1. (h *leadsHandler) => Nhận về từ file handler.go
// 2. ctx.Param("id") => Nhận ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *leadsHandler) FindByIDLead(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Tìm lead theo ID
	lead, err := h.service.FindByIDLead(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": lead})
}

// Hàm cập nhật lead
// Cấu trúc khai báo func (h *leadsHandler) UpdateLead(ctx *gin.Context) {
// 1. (h *leadsHandler) => Nhận về từ file handler.go
// 2. uuid.Parse(id) => Parse ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *leadsHandler) UpdateLead(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req LeadsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lead := &Lead{
		ID:         parsedID,
		Name:       &req.Name,
		Phone:      &req.Phone,
		Address:    &req.Address,
		Source:     &req.Source,
		Stage:      &req.Stage,
		Notes:      &req.Notes,
		UserID:     req.UserID,
		AssignedAt: req.AssignedAt,
		CampaignID: req.CampaignID,
		TenantID:   req.TenantID,
		ShowAt:     req.ShowAt,
		LastCallAt: req.LastCallAt,
	}

	// Cập nhật lead bằng service
	if err := h.service.UpdateLead(lead); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Lead updated successfully"})
}

// Hàm xóa lead
// Cấu trúc khai báo func (h *leadsHandler) DeleteLead(ctx *gin.Context) {
// 1. (h *leadsHandler) => Nhận về từ file handler.go
// 2. h.service.DeleteLead(id) => Gọi hàm xóa từ service
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *leadsHandler) DeleteLead(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Xóa lead bằng service
	if err := h.service.DeleteLead(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Lead deleted successfully"})
}
