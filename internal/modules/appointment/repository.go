package appointment

import "gorm.io/gorm"

// Định nghĩa các hàm sẽ sử dụng CRUD (Interface)
type AppointmentRepository interface {
	FindAllAppointments() ([]Appointment, error)         // (return về 1 array và error)
	FindByIDAppointment(id string) (*Appointment, error) // (return về 1 appointment và error)
	CreateAppointment(appointment *Appointment) error    // (return về error)
	UpdateAppointment(appointment *Appointment) error    // (return về error)
	DeleteAppointment(id string) error                   // (return về error)
}

// Định nghĩa các hàm sẽ sử dụng CRUD (Struct)
// Cấu trúc khai báo appointmentRepository struct {
// 1. db *gorm.DB => Nhận về từ file db.go
// 2. appointmentRepository => Trả về Struct
type appointmentRepository struct {
	db *gorm.DB
}

// Tạo mới repository
// Cấu trúc khai báo NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
// 1. (db *gorm.DB) => Nhận về từ file db.go
// 2. AppointmentRepository => Trả về Interface
// 3. return &appointmentRepository{db: db} => Trả về Struct
func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db} // (nhận về từ file db.go)
}

// Tìm tất cả appointments
// Cấu trúc khai báo FindAllAppointments() ([]Appointment, error) {
// 1. ([]Appointment, error) => Trả về 1 array và error
// 2. return &appointment{db: db} => Trả về Struct
func (r *appointmentRepository) FindAllAppointments() ([]Appointment, error) {
	var appointments []Appointment
	err := r.db.Find(&appointments).Error
	return appointments, err
}

// Tìm appointment theo ID
// Cấu trúc khai báo FindByIDAppointment(id string) (*Appointment, error) {
// 1. (id string) => Nhận về từ file handler.go
// 2. (*Appointment, error) => Trả về 1 appointment và error
// 3. return &appointment{db: db} => Trả về Struct
func (r *appointmentRepository) FindByIDAppointment(id string) (*Appointment, error) {
	var appointment Appointment
	err := r.db.Where("id = ?", id).First(&appointment).Error
	return &appointment, err
}

// Tạo mới appointment
// Cấu trúc khai báo CreateAppointment(appointment *Appointment) error {
// 1. (appointment *Appointment) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Create(appointment).Error => Trả về error
func (r *appointmentRepository) CreateAppointment(appointment *Appointment) error {
	return r.db.Create(appointment).Error // (return về error)
}

// Cập nhật appointment
// Cấu trúc khai báo UpdateAppointment(appointment *Appointment) error {
// 1. (appointment *Appointment) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Save(appointment).Error => Trả về error
func (r *appointmentRepository) UpdateAppointment(appointment *Appointment) error {
	return r.db.Save(appointment).Error // (return về error)
}

// Xóa appointment
// Cấu trúc khai báo DeleteAppointment(id string) error {
// 1. (id string) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return r.db.Where("id = ?", id).Delete(&Appointment{}).Error => Trả về error
func (r *appointmentRepository) DeleteAppointment(id string) error {
	return r.db.Where("id = ?", id).Delete(&Appointment{}).Error // (return về error)
}
