package campaigns

type CampaignsService interface {
	FindAllCampaigns() ([]Campaign, error)          // (return về 1 array và error)
	FindByIDCampaigns(id string) (*Campaign, error) // (return về 1 campaign và error)
	CreateCampaigns(campaigns *Campaign) error      // (return về error)
	UpdateCampaigns(campaigns *Campaign) error      // (return về error)
	DeleteCampaigns(id string) error                // (return về error)
}

// Struct
type campaignsService struct {
	campaignsRepository CampaignsRepository
}

// Tạo mới service
func NewCampaignsService(campaignsRepository CampaignsRepository) CampaignsService {
	return &campaignsService{campaignsRepository: campaignsRepository}
}

// Tìm tất cả Campaigns
func (s *campaignsService) FindAllCampaigns() ([]Campaign, error) {
	return s.campaignsRepository.FindAllCampaigns()
}

// Tìm Campaign theo ID
func (s *campaignsService) FindByIDCampaigns(id string) (*Campaign, error) {
	return s.campaignsRepository.FindByIDCampaigns(id)
}

// Tạo mới Campaign
func (s *campaignsService) CreateCampaigns(campaigns *Campaign) error {
	return s.campaignsRepository.CreateCampaigns(campaigns)
}

// Cập nhật Campaign
func (s *campaignsService) UpdateCampaigns(campaigns *Campaign) error {
	return s.campaignsRepository.UpdateCampaigns(campaigns)
}

// Xóa Campaign
func (s *campaignsService) DeleteCampaigns(id string) error {
	return s.campaignsRepository.DeleteCampaigns(id)
}
