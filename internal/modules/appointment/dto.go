package appointment

import (
	"errors"
	"time"
)

// Hàm nhập request client gửi lên
type AppointmentRequest struct {
	LeadID         string    `json:"leadId" binding:"required,uuid"`                                        // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	TenantID       string    `json:"tenantId" binding:"required,uuid"`                                      // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	CreatedByID    string    `json:"createdById" binding:"required,uuid"`                                   // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	Status         string    `json:"status" binding:"required,oneof=BOOKED CANCELLED NO_SHOW SHOWED"`                        // oneof=BOOKED CANCELLED NO_SHOW SHOWED khớp với database và model
	AppointmentAt  time.Time `json:"appointmentAt" binding:"required" time_format:"2006-01-02 15:04:05"`    // binding:"required" có nghĩa là bắt buộc phải có và phải là định dạng 2006-01-02 15:04:05
	Branch         string    `json:"branch" binding:"required,min=2,max=100"`                               // binding:"required,min=2,max=100" có nghĩa là bắt buộc phải có và phải có độ dài từ 2 đến 100 ký tự
	HasCompanion   bool      `json:"hasCompanion"`                                                          // Bool luôn có giá trị không cần phải binding
	CompanionName  *string   `json:"companionName,omitempty"`                                               // omitempty có nghĩa là không bắt buộc
	CompanionPhone *string   `json:"companionPhone,omitempty"`                                              // omitempty có nghĩa là không bắt buộc
	ShowDate       time.Time `json:"showDate" binding:"required" time_format:"2006-01-02"`                  // binding:"required" có nghĩa là bắt buộc phải có và phải là định dạng 2006-01-02
	IsClosed       bool      `json:"isClosed"`                                                              // Bool luôn có giá trị không cần phải binding
	ServiceType    string    `json:"serviceType" binding:"required,oneof=MEMBER PT YOGA"`                                   // oneof=MEMBER PT YOGA khớp với database và model
	Revenue        float64   `json:"revenue" binding:"required"`                                            // binding:"required" có nghĩa là bắt buộc phải có
}

// Hàm trả về cho client
type AppointmentResponse struct {
	ID             string    `json:"id"`
	LeadID         string    `json:"leadId"`
	TenantID       string    `json:"tenantId"`
	CreatedByID    string    `json:"createdById"`
	Status         string    `json:"status"`
	AppointmentAt  time.Time `json:"appointmentAt"`
	Branch         string    `json:"branch"`
	HasCompanion   bool      `json:"hasCompanion"`
	CompanionName  *string   `json:"companionName,omitempty"`
	CompanionPhone *string   `json:"companionPhone,omitempty"`
	ShowDate       time.Time `json:"showDate"`
	IsClosed       bool      `json:"isClosed"`
	ServiceType    string    `json:"serviceType"`
	Revenue        float64   `json:"revenue"`
}

// Hàm validate dữ liệu của hasCompanion
// Khi có hasCompanion = true thì companionName và companionPhone bắt buộc phải có
func (dto *AppointmentRequest) Validate() error {
	// Nếu HasCompanion = true thì kiểm tra companionName và companionPhone có tồn tại không
	if dto.HasCompanion {
		if dto.CompanionName == nil || dto.CompanionPhone == nil {
			return errors.New("companion info required") //  Trả về nếu không tìm thấy
		}
	}
	return nil // Trả về nil nếu không có lỗi
}
