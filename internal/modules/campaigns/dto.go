package campaigns

import "time"

// Hàm nhập request client gửi lên
type CampaignRequest struct {
	Name        *string `json:"name" binding:"omitempty,min=2,max=100"`        // omitempty có nghĩa là không bắt buộc
	Description *string `json:"description" binding:"omitempty,min=2,max=100"` // omitempty có nghĩa là không bắt buộc
	IsActive    *bool   `json:"isActive" binding:"omitempty"`                  // Bool luôn có giá trị không cần phải binding
	TenantID    *string `json:"tenantId" binding:"omitempty,uuid"`             // omitempty có nghĩa là không bắt buộc
}

// Hàm trả về cho client
type CampaignResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsActive    bool       `json:"isActive"`
	TenantID    string     `json:"tenantId"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
