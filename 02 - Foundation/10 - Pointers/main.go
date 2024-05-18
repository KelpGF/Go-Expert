package main

import "fmt"

func main() {
	// basic()
	// pointersAsParams()
	pointerAndStructs()
}

func basic() {
	a := 10
	fmt.Println("Value of a:", a)
	fmt.Println("Memory address of a:", &a)

	var pointer *int = &a
	fmt.Println("Value of pointer:", pointer)
	fmt.Println("Memory address of pointer:", &pointer)
	fmt.Println("Value of a using pointer:", *pointer)

	*pointer = 20
	fmt.Println("Value of a using pointer after change:", a)
	fmt.Println("Value of pointer:", pointer)
}

func pointersAsParams() {
	sum := func(a, b int) int {
		return a + b
	}
	sumPointer := func(a, b *int) int {
		*a = 100
		*b = 200
		return *a + *b
	}

	a := 10
	b := 20
	sumResult := sum(a, b)
	fmt.Println("A value:", a)
	fmt.Println("B value:", b)
	fmt.Println("Sum result:", sumResult)

	sumPointerResult := sumPointer(&a, &b)
	fmt.Println("A value:", a)
	fmt.Println("B value:", b)
	fmt.Println("Sum pointer result:", sumPointerResult)
}

type BankAccount struct {
	Balance int
}

// Looks like a constructor
func NewBankAccount() *BankAccount {
	return &BankAccount{Balance: 0}
}

// This method have a copy of the struct, so it will not change the original value
func (b BankAccount) simulateLoan(value int) int {
	b.Balance += value

	return b.Balance
}

// This method have a pointer to the struct, so it will change the original value
func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func pointerAndStructs() {
	account := NewBankAccount()
	fmt.Println("Initial balance:", account.Balance)

	account.Deposit(1000)
	fmt.Println("Deposit balance:", account.Balance)

	loan := account.simulateLoan(500)
	fmt.Println("Loan balance:", loan)
	fmt.Println("Original balance:", account.Balance)

	account.Deposit(100)
	fmt.Println("Deposit balance:", account.Balance)
}
