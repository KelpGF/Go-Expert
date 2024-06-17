package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Equal(t, 10.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 5.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	tax, err := CalculateTaxAndSave(1000, repository)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, tax)

	tax, err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, "error saving tax")
	assert.Equal(t, 0.0, tax)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
