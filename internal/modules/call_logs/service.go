package calllogs

type CallLogsService interface {
	FindAllCallLogs() ([]CallLog, error)          // (return về 1 array và error)
	FindByIDCallLogs(id string) (*CallLog, error) // (return về 1 calllog và error)
	CreateCallLogs(calllogs *CallLog) error       // (return về error)
	UpdateCallLogs(calllogs *CallLog) error       // (return về error)
	DeleteCallLogs(id string) error               // (return về error)
}

// Struct
type callLogsService struct {
	callLogsRepository CallLogsRepository
}

// Tạo mới service
func NewCallLogsService(callLogsRepository CallLogsRepository) CallLogsService {
	return &callLogsService{callLogsRepository: callLogsRepository}
}

// Tìm tất cả calllogs
func (s *callLogsService) FindAllCallLogs() ([]CallLog, error) {
	return s.callLogsRepository.FindAllCallLogs()
}

// Tìm calllog theo ID
func (s *callLogsService) FindByIDCallLogs(id string) (*CallLog, error) {
	return s.callLogsRepository.FindByIDCallLogs(id)
}

// Tạo mới calllog
func (s *callLogsService) CreateCallLogs(calllogs *CallLog) error {
	return s.callLogsRepository.CreateCallLogs(calllogs)
}

// Cập nhật calllog
func (s *callLogsService) UpdateCallLogs(calllogs *CallLog) error {
	return s.callLogsRepository.UpdateCallLogs(calllogs)
}

// Xóa calllog
func (s *callLogsService) DeleteCallLogs(id string) error {
	return s.callLogsRepository.DeleteCallLogs(id)
}
