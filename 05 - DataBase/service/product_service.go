package service

import (
	"database/sql"
	"gofc-database/model"
)

type productService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *productService {
	return &productService{db}
}

func (p *productService) Insert(product *model.Product) error {
	// prepare statement
	stmt, err := p.db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute statement
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) FindOne(id string) (*model.Product, error) {
	// prepare statement
	stmt, err := p.db.Prepare("SELECT id, name, price FROM products WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product model.Product

	// execute statement
	row := stmt.QueryRow(id)
	// scan row
	err = row.Scan(&product.ID, &product.Name, &product.Price)

	// same: err := stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productService) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *productService) Update(product *model.Product) error {
	// prepare statement
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute statement
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) Delete(id string) error {
	// prepare statement
	stmt, err := p.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute statement
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
