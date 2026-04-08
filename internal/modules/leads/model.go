package leads

import (
	"time"

	"github.com/google/uuid"
)

// Khai báo biến Stage
type Stage string

// Khai báo hằng số Stage
const (
	StageNew       Stage = "NEW"
	StageProcess   Stage = "PROCESS"
	StageContacted Stage = "CONTACTED"
	StageBooked    Stage = "BOOKED"
	StageShowed    Stage = "SHOWED"
	StageClosed    Stage = "CLOSED"
)

// Model lead
type Lead struct {
	ID         uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Name       *string    `gorm:"type:varchar(191)" json:"name"`
	Phone      *string    `gorm:"type:varchar(50);index" json:"phone"`
	Address    *string    `gorm:"type:varchar(191)" json:"address"`
	Source     *string    `gorm:"type:varchar(191)" json:"source"`
	Stage      *Stage     `gorm:"type:enum('NEW','PROCESS','CONTACTED','BOOKED','SHOWED','CLOSED');index" json:"stage,omitempty"`
	Notes      *string    `gorm:"type:varchar(191)" json:"notes"`
	UserID     *uuid.UUID `gorm:"type:char(36);index;column:userId" json:"userId"`
	AssignedAt *time.Time `gorm:"type:datetime" json:"assignedAt"`
	CampaignID *uuid.UUID `gorm:"type:char(36);index;column:campaignId" json:"campaignId"`
	TenantID   *uuid.UUID `gorm:"type:char(36);index;column:tenantId" json:"tenantId"`
	ShowAt     *time.Time `gorm:"type:datetime" json:"showAt"`
	LastCallAt *time.Time `gorm:"type:datetime;index" json:"lastCallAt"`

	CreatedAt *time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
