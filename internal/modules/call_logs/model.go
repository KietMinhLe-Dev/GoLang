package calllogs

import (
	"time"

	"github.com/google/uuid"
)

// Khai báo biến CallStatus
type CallStatus string

// Khai báo hằng số CallStatus
const (
	CallStatusConnected   CallStatus = "CONNECTED"
	CallStatusNoAnswer    CallStatus = "NO_ANSWER"
	CallStatusBusy        CallStatus = "BUSY"
	CallStatusUnreachable CallStatus = "UNREACHABLE"
	CallStatusWrongNumber CallStatus = "WRONG_NUMBER"
	CallStatusCancelled   CallStatus = "CANCELLED"
)

// Model call logs
type CallLog struct {
	ID           uuid.UUID   `gorm:"column:id; type:char(36);primaryKey" json:"id"`
	LeadID       *uuid.UUID  `gorm:"column:leadId; type:char(36)" json:"leadId"`
	UserID       *uuid.UUID  `gorm:"column:userId; type:char(36)" json:"userId"`
	TenantID     *uuid.UUID  `gorm:"column:tenantId; type:char(36)" json:"tenantId"`
	Status       *CallStatus `gorm:"column:status;type:enum('CONNECTED','NO_ANSWER','BUSY','UNREACHABLE','WRONG_NUMBER','CANCELLED')" json:"status,omitempty"`
	CancelReason *string     `gorm:"column:cancelReason; type:varchar(191)" json:"cancelReason"`
	Notes        *string     `gorm:"column:notes; type:varchar(191)" json:"notes"`
	CallID       *string     `gorm:"column:callid; type:varchar(191)" json:"callid"`
	Dnb          *string     `gorm:"column:dnb; type:varchar(191)" json:"dnb"`
	Ext          *string     `gorm:"column:ext; type:varchar(50)" json:"ext"`
	Phone        *string     `gorm:"column:phone; type:varchar(50)" json:"phone"`
	Direction    *string     `gorm:"column:direction; type:varchar(50)" json:"direction"`
	RecordingURL *string     `gorm:"column:recordingUrl; type:varchar(191)" json:"recordingUrl"`
	DurationSecs int         `gorm:"column:durationSecs" json:"durationSecs"`
	CalledAt     *time.Time  `gorm:"column:calledAt;type:timestamp" json:"calledAt,omitempty"`
}
