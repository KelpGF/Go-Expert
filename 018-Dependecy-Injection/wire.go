//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/KelpGF/Go-Expert/018-Dependecy-Injection/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository, // dependency factory
	wire.Bind(new(product.ProductRepository), new(*product.ProductRepositoryImpl)), // dependency binding
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}
