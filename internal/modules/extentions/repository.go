package extentions

import "gorm.io/gorm"

type ExtentionRepository interface {
	FindAllExtentions() ([]Extention, error)
	FindByIDExtention(id string) (*Extention, error)
	CreateExtention(extention *Extention) error
	UpdateExtention(extention *Extention) error
	DeleteExtention(id string) error
}

type extentionRepository struct {
	db *gorm.DB
}

func NewExtentionRepository(db *gorm.DB) ExtentionRepository {
	return &extentionRepository{db: db}
}

func (r *extentionRepository) FindAllExtentions() ([]Extention, error) {
	var extentions []Extention
	err := r.db.Find(&extentions).Error
	return extentions, err
}

func (r *extentionRepository) FindByIDExtention(id string) (*Extention, error) {
	var extention Extention
	err := r.db.Where("id = ?", id).First(&extention).Error
	return &extention, err
}

func (r *extentionRepository) CreateExtention(extention *Extention) error {
	return r.db.Create(extention).Error
}

func (r *extentionRepository) UpdateExtention(extention *Extention) error {
	return r.db.Save(extention).Error
}

func (r *extentionRepository) DeleteExtention(id string) error {
	return r.db.Where("id = ?", id).Delete(&Extention{}).Error
}
