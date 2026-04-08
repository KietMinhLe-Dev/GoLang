package notes

import (
	"time"

	"github.com/google/uuid"
)

// Note maps bảng notes (FK nullable theo csdl.sql).
type Note struct {
	ID          uuid.UUID  `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Content     *string    `gorm:"column:content;type:varchar(191)" json:"content,omitempty"`
	LeadID      *uuid.UUID `gorm:"column:leadId;type:char(36)" json:"leadId,omitempty"`
	CreatedByID *uuid.UUID `gorm:"column:createdById;type:char(36)" json:"createdById,omitempty"`
	TenantID    *uuid.UUID `gorm:"column:tenantId;type:char(36)" json:"tenantId,omitempty"`
	CreatedAt   *time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updatedAt" json:"updatedAt,omitempty"`
}
