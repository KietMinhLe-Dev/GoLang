package users

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	ACTIVE Status = "ACTIVE"
	BUSY   Status = "BUSY"
)

type Role string

const (
	ADMIN Role = "0"
	USER  Role = "1"
	AGENT Role = "2"
)

type User struct {
	ID       uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Email    *string    `gorm:"type:varchar(191);unique" json:"email"`
	Name     *string    `gorm:"type:varchar(191)" json:"name"`
	Status   *Status    `gorm:"type:enum('ACTIVE','BUSY')" json:"status"`
	Ext      *string    `gorm:"type:varchar(50)" json:"ext"`
	Role     *Role      `gorm:"type:enum('0','1','2')" json:"role"`
	TenantID *uuid.UUID `gorm:"type:char(36);index;column:tenantId" json:"tenantId"`

	CreatedAt *time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
