package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type UsersHandler interface {
	CreateUser(ctx *gin.Context)
	FindAllUsers(ctx *gin.Context)
	FindByIDUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo usersHandler struct {
// 1. service => Nhận về từ file service.go
// 2. usersHandler => Trả về Struct
type usersHandler struct {
	service UserService
}

// Tạo mới handler
// Cấu trúc khai báo NewUsersHandler(service UserService) UsersHandler {
// 1. (service UserService) => Nhận về từ file service.go
// 2. UsersHandler => Trả về Interface
// 3. return &usersHandler{service: service} => Trả về Struct
func NewUsersHandler(service UserService) UsersHandler {
	return &usersHandler{service: service}
}

// Hàm tạo mới user
// Cấu trúc khai báo func (h *usersHandler) CreateUser(ctx *gin.Context) {
// 1. (h *usersHandler) => Nhận về từ file handler.go
// 2. http.StatusCreated => Trả về mã lỗi 201
// 3. gin.H{} => Trả về JSON
func (h *usersHandler) CreateUser(ctx *gin.Context) {
	var req UserRequest // Khai báo biến req kiểu UserRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &User{
		ID:       uuid.New(),
		Email:    &req.Email,
		Name:     &req.Name,
		Status:   &req.Status,
		Ext:      &req.Ext,
		Role:     &req.Role,
		TenantID: &req.TenantID,
	}

	// Tạo mới user bằng service
	if err := h.service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": user})
}

// Hàm tìm tất cả users
// Cấu trúc khai báo func (h *usersHandler) FindAllUsers(ctx *gin.Context) {
// 1. (h *usersHandler) => Nhận về từ file handler.go
// 2. http.StatusOK => Trả về mã lỗi 200
// 3. gin.H{} => Trả về JSON
func (h *usersHandler) FindAllUsers(ctx *gin.Context) {
	// Tìm tất cả users
	users, err := h.service.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Hàm tìm user theo ID
// Cấu trúc khai báo func (h *usersHandler) FindByIDUser(ctx *gin.Context) {
// 1. (h *usersHandler) => Nhận về từ file handler.go
// 2. ctx.Param("id") => Nhận ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *usersHandler) FindByIDUser(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Tìm user theo ID
	user, err := h.service.FindByIDUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// Hàm cập nhật user
// Cấu trúc khai báo func (h *usersHandler) UpdateUser(ctx *gin.Context) {
// 1. (h *usersHandler) => Nhận về từ file handler.go
// 2. uuid.Parse(id) => Parse ID từ URL
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *usersHandler) UpdateUser(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &User{
		ID:       parsedID,
		Email:    &req.Email,
		Name:     &req.Name,
		Status:   &req.Status,
		Ext:      &req.Ext,
		Role:     &req.Role,
		TenantID: &req.TenantID,
	}

	// Cập nhật user bằng service
	if err := h.service.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Hàm xóa user
// Cấu trúc khai báo func (h *usersHandler) DeleteUser(ctx *gin.Context) {
// 1. (h *usersHandler) => Nhận về từ file handler.go
// 2. h.service.DeleteUser(id) => Gọi hàm xóa từ service
// 3. http.StatusOK => Trả về mã lỗi 200
func (h *usersHandler) DeleteUser(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Xóa user bằng service
	if err := h.service.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
