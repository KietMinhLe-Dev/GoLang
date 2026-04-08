package notes

import "time"

// Hàm nhận dữ liệu từ client
type NotesRequest struct {
	LeadID    string    `json:"leadId" binding:"required"`
	UserID    string    `json:"userId" binding:"required"`
	TenantID  string    `json:"tenantId" binding:"required"`
	Notes     string    `json:"notes" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
}

// Hàm trả về dữ liệu cho client
type NotesResponse struct {
	ID        string    `json:"id"`
	LeadID    string    `json:"leadId"`
	UserID    string    `json:"userId"`
	TenantID  string    `json:"tenantId"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


