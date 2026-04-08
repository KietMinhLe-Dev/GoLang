package campaigns

import (
	"time"

	"github.com/google/uuid"
)

// Campaign maps bảng campaigns (tenantId nullable theo csdl.sql).
type Campaign struct {
	ID          uuid.UUID  `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Name        *string    `gorm:"column:name;type:varchar(191)" json:"name,omitempty"`
	Description *string    `gorm:"column:description;type:varchar(191)" json:"description,omitempty"`
	IsActive    *bool      `gorm:"column:isActive;type:tinyint(1)" json:"isActive,omitempty"`
	TenantID    *uuid.UUID `gorm:"column:tenantId;type:char(36)" json:"tenantId,omitempty"`
	CreatedAt   *time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updatedAt" json:"updatedAt,omitempty"`
}
