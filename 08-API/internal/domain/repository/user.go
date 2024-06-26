package repository

import "github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
