package extentions

type ExtentionService interface {
	FindAllExtentions() ([]Extention, error)
	FindByIDExtention(id string) (*Extention, error)
	CreateExtention(extention *Extention) error
	UpdateExtention(extention *Extention) error
	DeleteExtention(id string) error
}

// Struct
type extentionService struct {
	extentionRepository ExtentionRepository
}

// Tạo mới service
func NewExtentionService(extentionRepository ExtentionRepository) ExtentionService {
	return &extentionService{extentionRepository: extentionRepository}
}

// Tìm tất cả extentions
func (s *extentionService) FindAllExtentions() ([]Extention, error) {
	return s.extentionRepository.FindAllExtentions()
}

// Tìm extention theo ID
func (s *extentionService) FindByIDExtention(id string) (*Extention, error) {
	return s.extentionRepository.FindByIDExtention(id)
}

// Tạo mới extention
func (s *extentionService) CreateExtention(extention *Extention) error {
	return s.extentionRepository.CreateExtention(extention)
}

// Cập nhật extention
func (s *extentionService) UpdateExtention(extention *Extention) error {
	return s.extentionRepository.UpdateExtention(extention)
}

// Xóa extention
func (s *extentionService) DeleteExtention(id string) error {
	return s.extentionRepository.DeleteExtention(id)
}
