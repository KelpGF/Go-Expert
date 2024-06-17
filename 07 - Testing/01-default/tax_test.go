package main

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500, 5},
		{1000, 10},
		{1500, 10},
		{0, 0},
	}

	for _, test := range tests {
		result := CalculateTax(test.amount)

		if result != test.expected {
			t.Errorf("Expected %f, but got %f", test.expected, result)
		}
	}
}

func FuzzCalculateTax(f *testing.F) {
	seeds := []float64{-100, -10, 0, 100, 1000, 1500, 10000, 30000}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Amount: %f. Expected 0, but got %f", amount, result)
		}

		if amount >= 20000 && result != 20 {
			t.Errorf("Amount: %f. Expected 20, but got %f", amount, result)
		}
	})
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}
