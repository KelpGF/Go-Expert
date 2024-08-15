package product

import "database/sql"

type ProductRepositoryTxTImpl struct {
	db *sql.DB
}

func NewProductRepositoryTxT(db *sql.DB) *ProductRepositoryTxTImpl {
	return &ProductRepositoryTxTImpl{db}
}

func (r ProductRepositoryTxTImpl) GetProduct(id int) (*Product, error) {
	return &Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
