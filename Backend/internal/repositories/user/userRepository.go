package repositories

import (
	"github.com/DevEmpulse/Stock-Management.git/internal/database"
	users "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *users.Users) error {
	result := database.DB.Create(user)
	return result.Error
}

func (r *UserRepository) FindUserByEmail(email string) (*users.Users, error) {
	var user users.Users
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindUserByID(id uint) (*users.Users, error) {
	var user users.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	result := database.DB.Delete(&users.Users{}, id)
	return result.Error
}

func (r *UserRepository) UpdateUser(id uint, updatedUser *users.Users) error {
	var user users.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(&user).Updates(updatedUser)
	return result.Error
}
