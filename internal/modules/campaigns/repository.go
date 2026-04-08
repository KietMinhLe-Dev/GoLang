package campaigns

import "gorm.io/gorm"

type CampaignsRepository interface {
	FindAllCampaigns() ([]Campaign, error)          // (return về 1 array và error)
	FindByIDCampaigns(id string) (*Campaign, error) // (return về 1 campaign và error)
	CreateCampaigns(campaigns *Campaign) error      // (return về error)
	UpdateCampaigns(campaigns *Campaign) error      // (return về error)
	DeleteCampaigns(id string) error                // (return về error)
}

// Định nghĩa các hàm sẽ sử dụng CRUD (Struct)
type campaignsRepository struct {
	db *gorm.DB
}

// Tạo mới repository
// Cấu trúc khai báo NewcampaignsRepository(db *gorm.DB) campaignsRepository {
// 1. (db *gorm.DB) => Nhận về từ file db.go
// 2. campaignsRepository => Trả về Interface
// 3. return &campaignsRepository{db: db} => Trả về Struct
func NewCampaignsRepository(db *gorm.DB) CampaignsRepository {
	return &campaignsRepository{db: db}
}

// Tìm tất cả Campaigns
// Cấu trúc khai báo FindAllCampaigns() ([]Campaign, error) {
// 1. ([]Campaign, error) => Trả về 1 array và error
// 2. return &campaignsRepository{db: db} => Trả về Struct
func (r *campaignsRepository) FindAllCampaigns() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Find(&campaigns).Error
	return campaigns, err
}

// Tìm Campaign theo ID
// Cấu trúc khai báo FindByIDCampaigns(id string) (*Campaign, error) {
// 1. (id string) => Nhận về từ file handler.go
// 2. (*Campaign, error) => Trả về 1 Campaign và error
// 3. return &campaignsRepository{db: db} => Trả về Struct
func (r *campaignsRepository) FindByIDCampaigns(id string) (*Campaign, error) {
	var Campaign Campaign
	err := r.db.Where("id = ?", id).First(&Campaign).Error
	return &Campaign, err
}

// Tạo mới Campaign
// Cấu trúc khai báo CreateCampaigns(Campaigns *Campaign) error {
// 1. (Campaigns *Campaign) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Create(Campaigns).Error => Trả về error
func (r *campaignsRepository) CreateCampaigns(Campaigns *Campaign) error {
	return r.db.Create(Campaigns).Error
}

// Cập nhật Campaign
// Cấu trúc khai báo UpdateCampaigns(Campaigns *Campaign) error {
// 1. (Campaigns *Campaign) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Save(Campaigns).Error => Trả về error
func (r *campaignsRepository) UpdateCampaigns(Campaigns *Campaign) error {
	return r.db.Save(Campaigns).Error
}

// Xóa Campaign
// Cấu trúc khai báo DeleteCampaigns(id string) error {
// 1. (id string) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Where("id = ?", id).Delete(&Campaign{}).Error => Trả về error
func (r *campaignsRepository) DeleteCampaigns(id string) error {
	return r.db.Where("id = ?", id).Delete(&Campaign{}).Error
}
