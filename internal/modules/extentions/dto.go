package extentions

type ExtentionRequest struct {
	Ext      string `json:"ext" binding:"required"`           // bắt buộc phải có
	Password string `json:"password" binding:"required"`      // bắt buộc phải có
	Name     string `json:"name" binding:"required"`          // bắt buộc phải có
	IsRecord bool   `json:"isRecord" binding:"required"`      // bắt buộc phải có
	TenantId string `json:"tenantId" binding:"required,uuid"` // bắt buộc phải có và là uuid
	UserId   string `json:"userId" binding:"required,uuid"`   // bắt buộc phải có và là uuid
}

type ExtentionResponse struct {
	ID       string `json:"id"`
	Ext      string `json:"ext"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsRecord bool   `json:"isRecord"`
	TenantId string `json:"tenantId"`
	UserId   string `json:"userId"`
}
