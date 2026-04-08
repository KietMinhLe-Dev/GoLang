package calllogs

import "time"

// Hàm nhập request client gửi lên
type CallLogsRequest struct {
	LeadID       string  `json:"leadId" binding:"required,uuid"`                                                              // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	UserID       string  `json:"userId" binding:"required,uuid"`                                                              // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	TenantID     string  `json:"tenantId" binding:"required,uuid"`                                                            // binding:"required,uuid" có nghĩa là bắt buộc phải có và phải là định dạng uuid
	Status       string  `json:"status" binding:"required,oneof=CONNECTED NO_ANSWER BUSY UNREACHABLE WRONG_NUMBER CANCELLED"` // oneof=CONNECTED NO_ANSWER BUSY UNREACHABLE WRONG_NUMBER CANCELLED có nghĩa là bắt buộc phải có và phải là một trong các giá trị trên
	CancelReason *string `json:"cancelReason,omitempty"`                                                                      // omitempty có nghĩa là không bắt buộc
	Notes        *string `json:"notes,omitempty"`                                                                             // omitempty có nghĩa là không bắt buộc
	CallID       *string `json:"callId,omitempty"`                                                                            // omitempty có nghĩa là không bắt buộc
	DNB          *string `json:"dnb,omitempty"`                                                                               // omitempty có nghĩa là không bắt buộc
	Ext          *string `json:"ext,omitempty"`                                                                               // omitempty có nghĩa là không bắt buộc
	Phone        *string `json:"phone,omitempty"`                                                                             // omitempty có nghĩa là không bắt buộc
	Direction    *string `json:"direction,omitempty"`                                                                         // omitempty có nghĩa là không bắt buộc
	RecordingURL *string `json:"recordingUrl,omitempty"`                                                                      // omitempty có nghĩa là không bắt buộc
	DurationSecs int     `json:"durationSecs"`                                                                                // binding:"required" có nghĩa là bắt buộc phải có
	CalledAt     string  `json:"calledAt" binding:"required" time_format:"2006-01-02 15:04:05"`                               // binding:"required" có nghĩa là bắt buộc phải có và phải là định dạng 2006-01-02 15:04:05
}

// Hàm trả về cho client
type CallLogsResponse struct {
	ID           string    `json:"id"`
	LeadID       string    `json:"leadId"`
	UserID       string    `json:"userId"`
	TenantID     string    `json:"tenantId"`
	Status       string    `json:"status"`
	CancelReason *string   `json:"cancelReason,omitempty"`
	Notes        *string   `json:"notes,omitempty"`
	CallID       *string   `json:"callId,omitempty"`
	DNB          *string   `json:"dnb,omitempty"`
	Ext          *string   `json:"ext,omitempty"`
	Phone        *string   `json:"phone,omitempty"`
	Direction    *string   `json:"direction,omitempty"`
	RecordingURL *string   `json:"recordingUrl,omitempty"`
	DurationSecs int       `json:"durationSecs"`
	CalledAt     time.Time `json:"calledAt"`
}
