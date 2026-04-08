package leads

import "gorm.io/gorm"

type LeadRepository interface {
	FindAllLeads() ([]Lead, error)         // (return về 1 array và error)
	FindByIDLead(id string) (*Lead, error) // (return về 1 lead và error)
	CreateLead(lead *Lead) error           // (return về error)
	UpdateLead(lead *Lead) error           // (return về error)
	DeleteLead(id string) error            // (return về error)
}

// Định nghĩa các hàm sẽ sử dụng CRUD (Struct)
type leadRepository struct {
	db *gorm.DB
}

// Tạo mới repository
// Cấu trúc khai báo NewLeadRepository(db *gorm.DB) LeadRepository {
// 1. (db *gorm.DB) => Nhận về từ file db.go
// 2. LeadRepository => Trả về Interface
// 3. return &leadRepository{db: db} => Trả về Struct
func NewLeadRepository(db *gorm.DB) LeadRepository {
	return &leadRepository{db: db}
}

// Tìm tất cả leads
// Cấu trúc khai báo FindAllLeads() ([]Lead, error) {
// 1. ([]Lead, error) => Trả về 1 array và error
// 2. return &leadRepository{db: db} => Trả về Struct
func (r *leadRepository) FindAllLeads() ([]Lead, error) {
	var leads []Lead
	err := r.db.Find(&leads).Error
	return leads, err
}

// Tìm lead theo ID
// Cấu trúc khai báo FindByIDLead(id string) (*Lead, error) {
// 1. (id string) => Nhận về từ file handler.go
// 2. (*Lead, error) => Trả về 1 lead và error
// 3. return &leadRepository{db: db} => Trả về Struct
func (r *leadRepository) FindByIDLead(id string) (*Lead, error) {
	var lead Lead
	err := r.db.Where("id = ?", id).First(&lead).Error
	return &lead, err
}

// Tạo mới lead
// Cấu trúc khai báo CreateLead(lead *Lead) error {
// 1. (lead *Lead) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Create(lead).Error => Trả về error
func (r *leadRepository) CreateLead(lead *Lead) error {
	return r.db.Create(lead).Error
}

// Cập nhật lead
// Cấu trúc khai báo UpdateLead(lead *Lead) error {
// 1. (lead *Lead) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Save(lead).Error => Trả về error
func (r *leadRepository) UpdateLead(lead *Lead) error {
	return r.db.Save(lead).Error
}

// Xóa lead
// Cấu trúc khai báo DeleteLead(id string) error {
// 1. (id string) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Where("id = ?", id).Delete(&Lead{}).Error => Trả về error
func (r *leadRepository) DeleteLead(id string) error {
	return r.db.Where("id = ?", id).Delete(&Lead{}).Error
}
