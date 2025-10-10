package main

import (
	"fmt"
	"strconv"
)

// TODO: Define the Account struct here
type Account struct {
	Name    string
	Balance float64
}

// TODO: Define the deposit method with pointer receiver here
func (a *Account) Deposit() {
	fmt.Printf("Account: %s, Balance: $%.2f\n", a.Name, a.Balance)
}

// TODO: Define the withdraw method with pointer receiver here
func (a *Account) Withdraw(transactionAmount float64) {
	fmt.Printf("Account: %s, Balance: $%.2f\n", a.Name, a.Balance+transactionAmount)
}

// TODO: Define the displayBalance method with value receiver here
func (a Account) displayBalance() {
	fmt.Printf("Account: %s, Balance: $%.2f\n", a.Name, a.Balance)
}
func main() {
	// Read input
	var name string
	var initialBalanceStr string
	var transactionAmountStr string

	fmt.Scanln(&name)
	fmt.Scanln(&initialBalanceStr)
	fmt.Scanln(&transactionAmountStr)

	// Convert string inputs to appropriate types
	initialBalance, _ := strconv.ParseFloat(initialBalanceStr, 64)
	transactionAmount, _ := strconv.ParseFloat(transactionAmountStr, 64)

	// TODO: Create an Account instance
	a := Account{
		Name:    name,
		Balance: initialBalance,
	}
	// TODO: Display initial balance, perform deposit, display balance, perform withdrawal, display final balance
	a.Deposit()
	a.Withdraw(transactionAmount)
	a.displayBalance()
}
