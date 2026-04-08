package tenants

import "time"

type TenantsRequest struct {
	Name          string `json:"name" binding:"required"`
	Status        string `json:"status" binding:"required"`
	MaxExtensions int    `json:"maxExtensions" binding:"required"`
	Extensions    int    `json:"extensions"`
	Domain        string `json:"domain"`
	KeyVoicecloud string `json:"keyVoicecloud"`
	PbxURL        string `json:"pbxUrl"`
	Timezone      string `json:"timezone"`
}

type TenantsResponse struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Status        string     `json:"status"`
	MaxExtensions int        `json:"maxExtensions"`
	Extensions    int        `json:"extensions"`
	Domain        string     `json:"domain"`
	KeyVoicecloud string     `json:"keyVoicecloud"`
	PbxURL        string     `json:"pbxUrl"`
	Timezone      string     `json:"timezone"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
