package repository

import (
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (userRepository *UserRepository) Create(user *entity.User) error {
	return userRepository.DB.Create(user).Error
}

func (userRepository *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := userRepository.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
