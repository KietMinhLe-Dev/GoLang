package extentions

import "github.com/google/uuid"

// Extention maps bảng extentions (tenantId, userId nullable theo csdl.sql).
type Extention struct {
	ID       uuid.UUID  `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Ext      *string    `gorm:"column:ext;type:varchar(50)" json:"ext,omitempty"`
	Password *string    `gorm:"column:password;type:varchar(191)" json:"password,omitempty"`
	Name     *string    `gorm:"column:name;type:varchar(191)" json:"name,omitempty"`
	IsRecord *bool      `gorm:"column:isRecord;type:tinyint(1)" json:"isRecord,omitempty"`
	TenantID *uuid.UUID `gorm:"column:tenantId;type:char(36)" json:"tenantId,omitempty"`
	UserID   *uuid.UUID `gorm:"column:userId;type:char(36)" json:"userId,omitempty"`
}
