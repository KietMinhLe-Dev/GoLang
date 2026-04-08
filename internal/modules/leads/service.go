package leads

type LeadService interface {
	FindAllLeads() ([]Lead, error)         // (return về 1 array và error)
	FindByIDLead(id string) (*Lead, error) // (return về 1 lead và error)
	CreateLead(lead *Lead) error           // (return về error)
	UpdateLead(lead *Lead) error           // (return về error)
	DeleteLead(id string) error            // (return về error)
}

// Struct
type leadService struct {
	leadRepository LeadRepository
}

// Tạo mới service
func NewLeadService(leadRepository LeadRepository) LeadService {
	return &leadService{leadRepository: leadRepository}
}

// Tìm tất cả leads
func (s *leadService) FindAllLeads() ([]Lead, error) {
	return s.leadRepository.FindAllLeads()
}

// Tìm lead theo ID
func (s *leadService) FindByIDLead(id string) (*Lead, error) {
	return s.leadRepository.FindByIDLead(id)
}

// Tạo mới lead
func (s *leadService) CreateLead(lead *Lead) error {
	return s.leadRepository.CreateLead(lead)
}

// Cập nhật lead
func (s *leadService) UpdateLead(lead *Lead) error {
	return s.leadRepository.UpdateLead(lead)
}

// Xóa lead
func (s *leadService) DeleteLead(id string) error {
	return s.leadRepository.DeleteLead(id)
}
