package users

import "gorm.io/gorm"

type UserRepository interface {
	FindAllUsers() ([]User, error)
	FindByIDUser(id string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByIDUser(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&User{}).Error
}
