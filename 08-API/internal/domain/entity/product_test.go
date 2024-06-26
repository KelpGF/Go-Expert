package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("product", 10.0)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "product", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotZero(t, product.CreatedAt)
}

func TestNewProductWhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("", 10.0)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductWhenPriceIsZero(t *testing.T) {
	product, err := NewProduct("product", 0)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestNewProductWhenPriceIsNegative(t *testing.T) {
	product, err := NewProduct("product", -10.0)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
