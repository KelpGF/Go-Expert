package product

import "database/sql"

type ProductRepository interface {
	GetProduct(id int) (*Product, error)
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (r ProductRepositoryImpl) GetProduct(id int) (*Product, error) {
	return &Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
