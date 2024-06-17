package main

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount should be greater than 0")
	}

	if amount >= 1000 && amount < 20000 {
		return 10, nil
	}

	if amount >= 20000 {
		return 20, nil
	}

	return 5, nil
}

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) (float64, error) {
	tax := CalculateTax2(amount)

	err := repository.SaveTax(tax)
	if err != nil {
		return 0, err
	}

	return tax, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0
	}

	return 5
}
