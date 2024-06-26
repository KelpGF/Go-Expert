package entity

import (
	"errors"
	"time"

	"github.com/KelpGF/Go-Expert/08-APIs/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("id is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price must be greater than 0")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(id entity.ID, name string, price float64) (*Product, error) {
	product := &Product{
		ID:        id,
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (product *Product) Validate() error {
	if product.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrInvalidID
	}

	if product.Name == "" {
		return ErrNameIsRequired
	}
	if product.Price == 0 {
		return ErrPriceIsRequired
	}
	if product.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
