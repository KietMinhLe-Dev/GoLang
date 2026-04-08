package appointment

import (
	"time"

	"github.com/google/uuid"
)

// AppointmentStatus khớp enum trong csdl.sql (appointments.status).
type AppointmentStatus string

const (
	AppointmentStatusBooked    AppointmentStatus = "BOOKED"
	AppointmentStatusCancelled AppointmentStatus = "CANCELLED"
	AppointmentStatusNoShow    AppointmentStatus = "NO_SHOW"
	AppointmentStatusShowed    AppointmentStatus = "SHOWED"
)

// AppointmentType khớp enum trong csdl.sql (appointments.serviceType).
type AppointmentType string

const (
	AppointmentTypeMember AppointmentType = "MEMBER"
	AppointmentTypePt     AppointmentType = "PT"
	AppointmentTypeYoga   AppointmentType = "YOGA"
)

// Appointment maps bảng appointments (các cột nullable dùng pointer).
type Appointment struct {
	ID             uuid.UUID          `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	LeadID         *uuid.UUID         `gorm:"column:leadId;type:char(36)" json:"leadId,omitempty"`
	TenantID       *uuid.UUID         `gorm:"column:tenantId;type:char(36)" json:"tenantId,omitempty"`
	CreatedByID    *uuid.UUID         `gorm:"column:createdById;type:char(36)" json:"createdById,omitempty"`
	Status         *AppointmentStatus `gorm:"column:status;type:enum('BOOKED','CANCELLED','NO_SHOW','SHOWED')" json:"status,omitempty"`
	AppointmentAt  *time.Time         `gorm:"column:appointmentAt;type:timestamp" json:"appointmentAt,omitempty"`
	Branch         *string            `gorm:"column:branch;type:varchar(191)" json:"branch,omitempty"`
	HasCompanion   *bool              `gorm:"column:hasCompanion;type:tinyint(1)" json:"hasCompanion,omitempty"`
	CompanionName  *string            `gorm:"column:companionName;type:varchar(191)" json:"companionName,omitempty"`
	CompanionPhone *string            `gorm:"column:companionPhone;type:varchar(50)" json:"companionPhone,omitempty"`
	ShowDate       *time.Time         `gorm:"column:showDate;type:timestamp" json:"showDate,omitempty"`
	IsClosed       *bool              `gorm:"column:isClosed;type:tinyint(1)" json:"isClosed,omitempty"`
	ServiceType    *AppointmentType   `gorm:"column:serviceType;type:enum('MEMBER','PT','YOGA')" json:"serviceType,omitempty"`
	Revenue        *float64           `gorm:"column:revenue;type:decimal(12,2)" json:"revenue,omitempty"`
}
