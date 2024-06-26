package repository

import "github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
