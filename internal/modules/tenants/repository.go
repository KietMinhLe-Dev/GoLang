package tenants

import "gorm.io/gorm"

type TenantRepository interface {
	FindAllTenants() ([]Tenant, error)
	FindByIDTenant(id string) (*Tenant, error)
	CreateTenant(tenant *Tenant) error
	UpdateTenant(tenant *Tenant) error
	DeleteTenant(id string) error
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) FindAllTenants() ([]Tenant, error) {
	var tenants []Tenant
	err := r.db.Find(&tenants).Error
	return tenants, err
}

func (r *tenantRepository) FindByIDTenant(id string) (*Tenant, error) {
	var tenant Tenant
	err := r.db.Where("id = ?", id).First(&tenant).Error
	return &tenant, err
}

func (r *tenantRepository) CreateTenant(tenant *Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *tenantRepository) UpdateTenant(tenant *Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *tenantRepository) DeleteTenant(id string) error {
	return r.db.Where("id = ?", id).Delete(&Tenant{}).Error
}
