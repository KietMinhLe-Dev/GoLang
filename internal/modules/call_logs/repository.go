package calllogs

import "gorm.io/gorm"

type CallLogsRepository interface {
	FindAllCallLogs() ([]CallLog, error)          // (return về 1 array và error)
	FindByIDCallLogs(id string) (*CallLog, error) // (return về 1 calllog và error)
	CreateCallLogs(calllogs *CallLog) error       // (return về error)
	UpdateCallLogs(calllogs *CallLog) error       // (return về error)
	DeleteCallLogs(id string) error               // (return về error)
}

// Định nghĩa các hàm sẽ sử dụng CRUD (Struct)
type callLogsRepository struct {
	db *gorm.DB
}

// Tạo mới repository
// Cấu trúc khai báo NewCallLogsRepository(db *gorm.DB) CallLogsRepository {
// 1. (db *gorm.DB) => Nhận về từ file db.go
// 2. CallLogsRepository => Trả về Interface
// 3. return &callLogsRepository{db: db} => Trả về Struct
func NewCallLogsRepository(db *gorm.DB) CallLogsRepository {
	return &callLogsRepository{db: db}
}

// Tìm tất cả calllogs
// Cấu trúc khai báo FindAllCallLogs() ([]CallLog, error) {
// 1. ([]CallLog, error) => Trả về 1 array và error
// 2. return &callLogsRepository{db: db} => Trả về Struct
func (r *callLogsRepository) FindAllCallLogs() ([]CallLog, error) {
	var calllogs []CallLog
	err := r.db.Find(&calllogs).Error
	return calllogs, err
}

// Tìm calllog theo ID
// Cấu trúc khai báo FindByIDCallLogs(id string) (*CallLog, error) {
// 1. (id string) => Nhận về từ file handler.go
// 2. (*CallLog, error) => Trả về 1 calllog và error
// 3. return &callLogsRepository{db: db} => Trả về Struct
func (r *callLogsRepository) FindByIDCallLogs(id string) (*CallLog, error) {
	var calllog CallLog
	err := r.db.Where("id = ?", id).First(&calllog).Error
	return &calllog, err
}

// Tạo mới calllog
// Cấu trúc khai báo CreateCallLogs(calllogs *CallLog) error {
// 1. (calllogs *CallLog) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Create(calllogs).Error => Trả về error
func (r *callLogsRepository) CreateCallLogs(calllogs *CallLog) error {
	return r.db.Create(calllogs).Error
}

// Cập nhật calllog
// Cấu trúc khai báo UpdateCallLogs(calllogs *CallLog) error {
// 1. (calllogs *CallLog) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Save(calllogs).Error => Trả về error
func (r *callLogsRepository) UpdateCallLogs(calllogs *CallLog) error {
	return r.db.Save(calllogs).Error
}

// Xóa calllog
// Cấu trúc khai báo DeleteCallLogs(id string) error {
// 1. (id string) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Where("id = ?", id).Delete(&CallLog{}).Error => Trả về error
func (r *callLogsRepository) DeleteCallLogs(id string) error {
	return r.db.Where("id = ?", id).Delete(&CallLog{}).Error
}
