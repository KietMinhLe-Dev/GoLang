package tenants

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type TenantsHandler interface {
	CreateTenant(ctx *gin.Context)
	FindAllTenants(ctx *gin.Context)
	FindByIDTenant(ctx *gin.Context)
	UpdateTenant(ctx *gin.Context)
	DeleteTenant(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo tenantsHandler struct {
// 1. service => Nhận về từ file service.go
// 2. tenantsHandler => Trả về Struct
type tenantsHandler struct {
	service TenantService
}

// Tạo mới handler
// Cấu trúc khai báo NewTenantsHandler(service TenantService) TenantsHandler {
// 1. (service TenantService) => Nhận về từ file service.go
// 2. TenantsHandler => Trả về Interface
// 3. return &tenantsHandler{service: service} => Trả về Struct
func NewTenantsHandler(service TenantService) TenantsHandler {
	return &tenantsHandler{service: service}
}

// Hàm tạo mới tenant
// Cấu trúc khai báo func (h *tenantsHandler) CreateTenant(ctx *gin.Context) {
// 1. (h *tenantsHandler) => Nhận về từ file handler.go
// 2. http.StatusCreated => Trả về mã lỗi 201
// 3. gin.H{} => Trả về JSON
func (h *tenantsHandler) CreateTenant(ctx *gin.Context) {
	var req TenantsRequest // Khai báo biến req kiểu TenantsRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant := &Tenant{
		ID:            uuid.New(),
		Name:          &req.Name,
		Status:        &req.Status,
		MaxExtensions: &req.MaxExtensions,
		Extensions:    &req.Extensions,
		Domain:        &req.Domain,
		KeyVoicecloud: &req.KeyVoicecloud,
		PbxURL:        &req.PbxURL,
		Timezone:      &req.Timezone,
	}

	// Tạo mới tenant bằng service
	if err := h.service.CreateTenant(tenant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusCreated, gin.H{"message": "Tenant created successfully", "data": tenant})
}

// Hàm tìm tất cả tenants
// Cấu trúc khai báo func (h *tenantsHandler) FindAllTenants(ctx *gin.Context) {
// 1. (h *tenantsHandler) => Nhận về từ file handler.go
// 2. http.StatusOK => Trả về mã lỗi 200
// 3. gin.H{} => Trả về JSON
func (h *tenantsHandler) FindAllTenants(ctx *gin.Context) {
	// Tìm tất cả tenants
	tenants, err := h.service.FindAllTenants()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": tenants})
}

// Hàm tìm tenant theo ID
// Cấu trúc khai báo func (h *tenantsHandler) FindByIDTenant(ctx *gin.Context) {
// 1. (h *tenantsHandler) => Nhận về từ file handler.go
// 2. ctx.Param("id") => Nhận ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *tenantsHandler) FindByIDTenant(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Tìm tenant theo ID
	tenant, err := h.service.FindByIDTenant(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": tenant})
}

// Hàm cập nhật tenant
// Cấu trúc khai báo func (h *tenantsHandler) UpdateTenant(ctx *gin.Context) {
// 1. (h *tenantsHandler) => Nhận về từ file handler.go
// 2. uuid.Parse(id) => Parse ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *tenantsHandler) UpdateTenant(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req TenantsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant := &Tenant{
		ID:            parsedID,
		Name:          &req.Name,
		Status:        &req.Status,
		MaxExtensions: &req.MaxExtensions,
		Extensions:    &req.Extensions,
		Domain:        &req.Domain,
		KeyVoicecloud: &req.KeyVoicecloud,
		PbxURL:        &req.PbxURL,
		Timezone:      &req.Timezone,
	}

	// Cập nhật tenant bằng service
	if err := h.service.UpdateTenant(tenant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Tenant updated successfully"})
}

// Hàm xóa tenant
// Cấu trúc khai báo func (h *tenantsHandler) DeleteTenant(ctx *gin.Context) {
// 1. (h *tenantsHandler) => Nhận về từ file handler.go
// 2. h.service.DeleteTenant(id) => Gọi hàm xóa từ service
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *tenantsHandler) DeleteTenant(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Xóa tenant bằng service
	if err := h.service.DeleteTenant(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}
