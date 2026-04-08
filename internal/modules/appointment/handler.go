package appointment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type AppointmentHandler interface {
	CreateAppointment(ctx *gin.Context)
	FindAllAppointments(ctx *gin.Context)
	FindByIDAppointment(ctx *gin.Context)
	UpdateAppointment(ctx *gin.Context)
	DeleteAppointment(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo appointmentHandler struct {
// 1. service => Nhận về từ file service.go
// 2. appointmentHandler => Trả về Struct
type appointmentHandler struct {
	service AppointmentService
}

// Tạo mới handler
// Cấu trúc khai báo NewAppointmentHandler(service AppointmentService) AppointmentHandler {
// 1. (service AppointmentService) => Nhận về từ file service.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func NewAppointmentHandler(service AppointmentService) AppointmentHandler {
	return &appointmentHandler{service: service}
}

// Hàm tạo mới appointment
// Cấu trúc khai báo func (h *appointmentHandler) CreateAppointment(ctx *gin.Context) {
// 1. (h *appointmentHandler) => Nhận về từ file handler.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func (h *appointmentHandler) CreateAppointment(ctx *gin.Context) {
	var req AppointmentRequest // Khai báo biến req kiểu AppointmentRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	// Cấu trúc khai báo if err := ctx.ShouldBindJSON(&req); err != nil {
	// 1. ctx.ShouldBindJSON(&req) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate logic nâng cao (ví dụ: Companion info)
	// Cấu trúc khai báo if err := req.Validate(); err != nil {
	// 1. req.Validate() => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi từ DTO sang Model
	// Cấu trúc khai báo leadID, _ := uuid.Parse(req.LeadID)
	// 1. uuid.Parse(req.LeadID) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	leadID, _ := uuid.Parse(req.LeadID)
	tenantID, _ := uuid.Parse(req.TenantID)
	createdByID, _ := uuid.Parse(req.CreatedByID)
	status := AppointmentStatus(req.Status)
	serviceType := AppointmentType(req.ServiceType)

	// Tạo mới appointment
	// Cấu trúc khai báo appointment := &Appointment{
	// 1. appointment := &Appointment{} => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	appointment := &Appointment{
		ID:             uuid.New(), // Tạo UUID mới cho bản ghi mới
		LeadID:         &leadID,
		TenantID:       &tenantID,
		CreatedByID:    &createdByID,
		Status:         &status,
		AppointmentAt:  &req.AppointmentAt,
		Branch:         &req.Branch,
		HasCompanion:   &req.HasCompanion,
		CompanionName:  req.CompanionName,
		CompanionPhone: req.CompanionPhone,
		ShowDate:       &req.ShowDate,
		IsClosed:       &req.IsClosed,
		ServiceType:    &serviceType,
		Revenue:        &req.Revenue,
	}

	// Tạo mới appointment
	// Cấu trúc khai báo if err := h.service.CreateAppointment(appointment); err != nil {
	// 1. h.service.CreateAppointment(appointment) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	if err := h.service.CreateAppointment(appointment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	// Cấu trúc khai báo ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "data": appointment})
	// 1. ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "data": appointment}) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "data": appointment})
}

// Hàm tìm tất cả appointment
// Cấu trúc khai báo func (h *appointmentHandler) FindAllAppointments(ctx *gin.Context) {
// 1. (h *appointmentHandler) => Nhận về từ file handler.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func (h *appointmentHandler) FindAllAppointments(ctx *gin.Context) {
	// Tìm tất cả appointment
	// Cấu trúc khai báo appointments, err := h.service.FindAllAppointments()
	// 1. appointments, err := h.service.FindAllAppointments() => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	appointments, err := h.service.FindAllAppointments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về
	ctx.JSON(http.StatusOK, gin.H{"data": appointments})
}

// Hàm tìm appointment theo ID
// Cấu trúc khai báo func (h *appointmentHandler) FindByIDAppointment(ctx *gin.Context) {
// 1. (h *appointmentHandler) => Nhận về từ file handler.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func (h *appointmentHandler) FindByIDAppointment(ctx *gin.Context) {
	// Lấy ID từ URL
	// Cấu trúc khai báo id := ctx.Param("id")
	// 1. id := ctx.Param("id") => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	id := ctx.Param("id")
	// Tìm appointment theo ID
	// Cấu trúc khai báo appointment, err := h.service.FindByIDAppointment(id)
	// 1. appointment, err := h.service.FindByIDAppointment(id) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	appointment, err := h.service.FindByIDAppointment(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Trả về
	ctx.JSON(http.StatusOK, gin.H{"data": appointment})
}

// Hàm cập nhật appointment
// Cấu trúc khai báo func (h *appointmentHandler) UpdateAppointment(ctx *gin.Context) {
// 1. (h *appointmentHandler) => Nhận về từ file handler.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func (h *appointmentHandler) UpdateAppointment(ctx *gin.Context) {
	// Lấy ID từ URL
	// Cấu trúc khai báo id := ctx.Param("id")
	// 1. id := ctx.Param("id") => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	id := ctx.Param("id")
	// Parse ID
	// Cấu trúc khai báo parsedID, err := uuid.Parse(id)
	// 1. parsedID, err := uuid.Parse(id) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Khai báo biến req kiểu AppointmentRequest
	// Cấu trúc khai báo var req AppointmentRequest
	// 1. var req AppointmentRequest => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	var req AppointmentRequest
	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	// Cấu trúc khai báo if err := ctx.ShouldBindJSON(&req); err != nil {
	// 1. ctx.ShouldBindJSON(&req) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse ID
	// Cấu trúc khai báo leadID, _ := uuid.Parse(req.LeadID)
	// 1. leadID, _ := uuid.Parse(req.LeadID) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	leadID, _ := uuid.Parse(req.LeadID)
	tenantID, _ := uuid.Parse(req.TenantID)
	createdByID, _ := uuid.Parse(req.CreatedByID)
	status := AppointmentStatus(req.Status)
	serviceType := AppointmentType(req.ServiceType)

	// Tạo mới appointment
	// Cấu trúc khai báo appointment := &Appointment{
	// 1. appointment := &Appointment{} => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	appointment := &Appointment{
		ID:             parsedID,
		LeadID:         &leadID,
		TenantID:       &tenantID,
		CreatedByID:    &createdByID,
		Status:         &status,
		AppointmentAt:  &req.AppointmentAt,
		Branch:         &req.Branch,
		HasCompanion:   &req.HasCompanion,
		CompanionName:  req.CompanionName,
		CompanionPhone: req.CompanionPhone,
		ShowDate:       &req.ShowDate,
		IsClosed:       &req.IsClosed,
		ServiceType:    &serviceType,
		Revenue:        &req.Revenue,
	}

	// Cập nhật appointment
	// Cấu trúc khai báo if err := h.service.UpdateAppointment(appointment); err != nil {
	// 1. h.service.UpdateAppointment(appointment) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	if err := h.service.UpdateAppointment(appointment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	// Cấu trúc khai báo ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
	// 1. ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"}) => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

// Hàm xóa appointment
// Cấu trúc khai báo func (h *appointmentHandler) DeleteAppointment(ctx *gin.Context) {
// 1. (h *appointmentHandler) => Nhận về từ file handler.go
// 2. AppointmentHandler => Trả về Interface
// 3. return &appointmentHandler{service: service} => Trả về Struct
func (h *appointmentHandler) DeleteAppointment(ctx *gin.Context) {
	// Lấy ID từ URL
	// Cấu trúc khai báo id := ctx.Param("id")
	// 1. id := ctx.Param("id") => Nhận về từ file handler.go
	// 2. AppointmentHandler => Trả về Interface
	// 3. return &appointmentHandler{service: service} => Trả về Struct
	id := ctx.Param("id")
	if err := h.service.DeleteAppointment(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
