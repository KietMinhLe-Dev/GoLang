package appointment

type AppointmentService interface {
	FindAllAppointments() ([]Appointment, error)         // (return về 1 array và error)
	FindByIDAppointment(id string) (*Appointment, error) // (return về 1 appointment và error)
	CreateAppointment(appointment *Appointment) error    // (return về error)
	UpdateAppointment(appointment *Appointment) error    // (return về error)
	DeleteAppointment(id string) error                   // (return về error)
}

// Struct
// Cấu trúc khai báo appointmentService struct {
// 1. appointmentRepository => Nhận về từ file repository.go
// 2. appointmentService => Trả về Struct
type appointmentService struct {
	appointmentRepository AppointmentRepository
}

// Tạo mới service
// Cấu trúc khai báo NewAppointmentService(appointmentRepository AppointmentRepository) AppointmentService {
// 1. (appointmentRepository AppointmentRepository) => Nhận về từ file repository.go
// 2. AppointmentService => Trả về Interface
// 3. return &appointmentService{appointmentRepository: appointmentRepository} => Trả về Struct
func NewAppointmentService(appointmentRepository AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepository: appointmentRepository}
}

// Tìm tất cả appointments
// Cấu trúc khai báo FindAllAppointments() ([]Appointment, error) {
// 1. ([]Appointment, error) => Trả về 1 array và error
// 2. return &appointmentService{appointmentRepository: appointmentRepository} => Trả về Struct
func (a *appointmentService) FindAllAppointments() ([]Appointment, error) {
	return a.appointmentRepository.FindAllAppointments()
}

// Tìm appointment theo ID
// Cấu trúc khai báo FindByIDAppointment(id string) (*Appointment, error) {
// 1. (id string) => Nhận về từ file handler.go
// 2. (*Appointment, error) => Trả về 1 appointment và error
// 3. return &appointmentService{appointmentRepository: appointmentRepository} => Trả về Struct
func (a *appointmentService) FindByIDAppointment(id string) (*Appointment, error) {
	return a.appointmentRepository.FindByIDAppointment(id)
}

// Tạo mới appointment
// Cấu trúc khai báo CreateAppointment(appointment *Appointment) error {
// 1. (appointment *Appointment) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return a.appointmentRepository.CreateAppointment(appointment) => Trả về error
func (a *appointmentService) CreateAppointment(appointment *Appointment) error {
	return a.appointmentRepository.CreateAppointment(appointment)
}

// Cập nhật appointment
// Cấu trúc khai báo UpdateAppointment(appointment *Appointment) error {
// 1. (appointment *Appointment) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return a.appointmentRepository.UpdateAppointment(appointment) => Trả về error
func (a *appointmentService) UpdateAppointment(appointment *Appointment) error {
	return a.appointmentRepository.UpdateAppointment(appointment)
}

// Xóa appointment
// Cấu trúc khai báo DeleteAppointment(id string) error {
// 1. (id string) => Nhận về từ file handler.go
// 2. error => Trả về error
// 3. return a.appointmentRepository.DeleteAppointment(id) => Trả về error
func (a *appointmentService) DeleteAppointment(id string) error {
	return a.appointmentRepository.DeleteAppointment(id)
}
