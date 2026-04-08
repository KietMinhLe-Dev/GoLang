package tenants

import (
	"time"

	"github.com/google/uuid"
)

// Model tenants
type Tenant struct {
	ID            uuid.UUID  `gorm:"column:id; type:char(36);primaryKey" json:"id"`
	Name          *string    `gorm:"column:name; type:varchar(191)" json:"name"`
	Status        *string    `gorm:"column:status; type:varchar(191)" json:"status"`
	MaxExtensions *int       `gorm:"column:max_extensions; type:int" json:"maxExtensions"`
	Extensions    *int       `gorm:"column:extensions; type:int" json:"extensions"`
	Domain        *string    `gorm:"column:domain; type:varchar(191)" json:"domain"`
	KeyVoicecloud *string    `gorm:"column:key_voicecloud; type:varchar(191)" json:"keyVoicecloud"`
	PbxURL        *string    `gorm:"column:pbx_url; type:varchar(191)" json:"pbxUrl"`
	Timezone      *string    `gorm:"column:timezone; type:varchar(191)" json:"timezone"`
	CreatedAt     *time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
