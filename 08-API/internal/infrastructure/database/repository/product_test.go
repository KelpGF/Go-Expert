package repository

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/KelpGF/Go-Expert/08-APIs/test/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, productRepository := makeSut()

	product := defaultProduct()
	err := productRepository.Create(product)

	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, productRepository := makeSut()

	for i := 1; i <= 25; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		db.Create(product)
	}

	products, err := productRepository.FindAll(0, 0, "")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productRepository.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productRepository.FindAll(2, 10, "desc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 15", products[0].Name)
	assert.Equal(t, "Product 6", products[9].Name)
}

func TestFindByID(t *testing.T) {
	db, productRepository := makeSut()

	product := defaultProduct()
	db.Create(product)

	productFound, err := productRepository.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, productRepository := makeSut()

	product := defaultProduct()
	db.Create(product)

	product.Name = "new name"
	product.Price = 20.5

	err := productRepository.Update(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, productRepository := makeSut()

	product := defaultProduct()
	db.Create(product)

	err := productRepository.Delete(product.ID.String())
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, product.ID).Error

	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Empty(t, productFound)
}

func makeSut() (*gorm.DB, *ProductRepository) {
	db := setupProduct()
	productRepository := NewProductRepository(db)

	return db, productRepository
}

func setupProduct() *gorm.DB {
	db := database.Setup()

	db.AutoMigrate(&entity.Product{})

	return db
}

func defaultProduct() *entity.Product {
	product, _ := entity.NewProduct("product", 10.5)

	return product
}
