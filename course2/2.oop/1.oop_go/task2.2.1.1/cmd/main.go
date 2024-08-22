package main

import (
	"fmt"

	"task2.2.1.1/accounts"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

func ProcessAccount(a Accounter) {
	a.Deposit(400)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &accounts.CurrentAccount{}
	s := &accounts.SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}
