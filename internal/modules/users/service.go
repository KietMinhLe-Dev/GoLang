package users

type UserService interface {
	FindAllUsers() ([]User, error)
	FindByIDUser(id string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

// Struct
type userService struct {
	userRepository UserRepository
}

// Tạo mới service
func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

// Tìm tất cả users
func (s *userService) FindAllUsers() ([]User, error) {
	return s.userRepository.FindAllUsers()
}

// Tìm user theo ID
func (s *userService) FindByIDUser(id string) (*User, error) {
	return s.userRepository.FindByIDUser(id)
}

// Tạo mới user
func (s *userService) CreateUser(user *User) error {
	return s.userRepository.CreateUser(user)
}

// Cập nhật user
func (s *userService) UpdateUser(user *User) error {
	return s.userRepository.UpdateUser(user)
}

// Xóa user
func (s *userService) DeleteUser(id string) error {
	return s.userRepository.DeleteUser(id)
}
