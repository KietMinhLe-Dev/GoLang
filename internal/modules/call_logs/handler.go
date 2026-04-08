package calllogs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Interface
type CallLogsHandler interface {
	CreateCallLogs(ctx *gin.Context)
	FindAllCallLogs(ctx *gin.Context)
	FindByIDCallLogs(ctx *gin.Context)
	UpdateCallLogs(ctx *gin.Context)
	DeleteCallLogs(ctx *gin.Context)
}

// Struct
// Cấu trúc khai báo callLogsHandler struct {
// 1. service => Nhận về từ file service.go
// 2. callLogsHandler => Trả về Struct
type callLogsHandler struct {
	service CallLogsService
}

// Tạo mới handler
// Cấu trúc khai báo NewCallLogsHandler(service CallLogsService) CallLogsHandler {
// 1. (service CallLogsService) => Nhận về từ file service.go
// 2. CallLogsHandler => Trả về Interface
// 3. return &callLogsHandler{service: service} => Trả về Struct
func NewCallLogsHandler(service CallLogsService) CallLogsHandler {
	return &callLogsHandler{service: service}
}

// Hàm tạo mới calllog
// Cấu trúc khai báo func (h *callLogsHandler) CreateCallLogs(ctx *gin.Context) {
// 1. (h *callLogsHandler) => Nhận về từ file handler.go
// 2. http.StatusCreated => Trả về mã lỗi 201
// 3. gin.H{} => Trả về JSON
func (h *callLogsHandler) CreateCallLogs(ctx *gin.Context) {
	var req CallLogsRequest // Khai báo biến req kiểu CallLogsRequest

	// Kiểm tra lỗi khi nhận dữ liệu từ Client
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi dữ liệu
	leadID, _ := uuid.Parse(req.LeadID)
	userID, _ := uuid.Parse(req.UserID)
	tenantID, _ := uuid.Parse(req.TenantID)
	status := CallStatus(req.Status)
	calledAt, _ := time.Parse("2006-01-02 15:04:05", req.CalledAt)

	calllog := &CallLog{
		ID:           uuid.New(),
		LeadID:       &leadID,
		UserID:       &userID,
		TenantID:     &tenantID,
		Status:       &status,
		CancelReason: req.CancelReason,
		Notes:        req.Notes,
		CallID:       req.CallID,
		Dnb:          req.DNB,
		Ext:          req.Ext,
		Phone:        req.Phone,
		Direction:    req.Direction,
		RecordingURL: req.RecordingURL,
		DurationSecs: req.DurationSecs,
		CalledAt:     &calledAt,
	}

	// Tạo mới calllog bằng service
	if err := h.service.CreateCallLogs(calllog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusCreated, gin.H{"message": "Call log created successfully", "data": calllog})
}

// Hàm tìm tất cả calllogs
// Cấu trúc khai báo func (h *callLogsHandler) FindAllCallLogs(ctx *gin.Context) {
// 1. (h *callLogsHandler) => Nhận về từ file handler.go
// 2. http.StatusOK => Trả về mã lỗi 200
// 3. gin.H{} => Trả về JSON
func (h *callLogsHandler) FindAllCallLogs(ctx *gin.Context) {
	// Tìm tất cả calllogs
	calllogs, err := h.service.FindAllCallLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": calllogs})
}

// Hàm tìm calllog theo ID
// Cấu trúc khai báo func (h *callLogsHandler) FindByIDCallLogs(ctx *gin.Context) {
// 1. (h *callLogsHandler) => Nhận về từ file handler.go
// 2. ctx.Param("id") => Nhận ID từ URL
func (h *callLogsHandler) FindByIDCallLogs(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Tìm calllog theo ID
	calllog, err := h.service.FindByIDCallLogs(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Call log not found"})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"data": calllog})
}

// Hàm cập nhật calllog
// Cấu trúc khai báo func (h *callLogsHandler) UpdateCallLogs(ctx *gin.Context) {
// 1. (h *callLogsHandler) => Nhận về từ file handler.go
// 2. uuid.Parse(id) => Parse ID từ URL
func (h *callLogsHandler) UpdateCallLogs(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req CallLogsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leadID, _ := uuid.Parse(req.LeadID)
	userID, _ := uuid.Parse(req.UserID)
	tenantID, _ := uuid.Parse(req.TenantID)
	status := CallStatus(req.Status)
	calledAt, _ := time.Parse("2006-01-02 15:04:05", req.CalledAt)

	calllog := &CallLog{
		ID:           parsedID,
		LeadID:       &leadID,
		UserID:       &userID,
		TenantID:     &tenantID,
		Status:       &status,
		CancelReason: req.CancelReason,
		Notes:        req.Notes,
		CallID:       req.CallID,
		Dnb:          req.DNB,
		Ext:          req.Ext,
		Phone:        req.Phone,
		Direction:    req.Direction,
		RecordingURL: req.RecordingURL,
		DurationSecs: req.DurationSecs,
		CalledAt:     &calledAt,
	}

	// Cập nhật calllog bằng service
	if err := h.service.UpdateCallLogs(calllog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Call log updated successfully"})
}

// Hàm xóa calllog
// Cấu trúc khai báo func (h *callLogsHandler) DeleteCallLogs(ctx *gin.Context) {
// 1. (h *callLogsHandler) => Nhận về từ file handler.go
func (h *callLogsHandler) DeleteCallLogs(ctx *gin.Context) {
	// Lấy ID từ URL
	id := ctx.Param("id")

	// Xóa calllog bằng service
	if err := h.service.DeleteCallLogs(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{"message": "Call log deleted successfully"})
}
