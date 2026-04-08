package tenants

type TenantService interface {
	FindAllTenants() ([]Tenant, error)
	FindByIDTenant(id string) (*Tenant, error)
	CreateTenant(tenant *Tenant) error
	UpdateTenant(tenant *Tenant) error
	DeleteTenant(id string) error
}

// Struct
type tenantService struct {
	tenantRepository TenantRepository
}

// Tạo mới service
func NewTenantService(tenantRepository TenantRepository) TenantService {
	return &tenantService{tenantRepository: tenantRepository}
}

// Tìm tất cả tenants
func (s *tenantService) FindAllTenants() ([]Tenant, error) {
	return s.tenantRepository.FindAllTenants()
}

// Tìm tenant theo ID
func (s *tenantService) FindByIDTenant(id string) (*Tenant, error) {
	return s.tenantRepository.FindByIDTenant(id)
}

// Tạo mới tenant
func (s *tenantService) CreateTenant(tenant *Tenant) error {
	return s.tenantRepository.CreateTenant(tenant)
}

// Cập nhật tenant
func (s *tenantService) UpdateTenant(tenant *Tenant) error {
	return s.tenantRepository.UpdateTenant(tenant)
}

// Xóa tenant
func (s *tenantService) DeleteTenant(id string) error {
	return s.tenantRepository.DeleteTenant(id)
}
