package main

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func NewCustomer(id int, opts ...CustomerOption) *Customer {
	customer := &Customer{
		ID: id,
	}

	for _, opt := range opts {
		opt(customer)
	}

	return customer
}

type SavingsAccount struct {
	mu      sync.RWMutex
	balance float64
}

func (sa *SavingsAccount) Balance() float64 {
	sa.mu.RLock()
	balance := sa.balance
	sa.mu.RUnlock()
	return balance
}

func (sa *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	sa.mu.Lock()
	sa.balance += amount
	sa.mu.Unlock()
	return nil
}

func (sa *SavingsAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	sa.mu.Lock()
	if sa.balance < 1000 {
		return fmt.Errorf("your current balance is lower than 1000")
	}
	if sa.balance < amount {
		return fmt.Errorf("not enough mony")
	}
	sa.balance -= amount
	sa.mu.Unlock()
	return nil
}

type CheckingAccount struct {
	mu      sync.RWMutex
	balance float64
}

func (ca *CheckingAccount) Balance() float64 {
	ca.mu.RLock()
	balance := ca.balance
	ca.mu.RUnlock()
	return balance
}

func (ca *CheckingAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	ca.mu.Lock()
	ca.balance += amount
	ca.mu.Unlock()
	return nil
}

func (ca *CheckingAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	ca.mu.Lock()
	if ca.balance < amount {
		return fmt.Errorf("not enough mony")
	}
	ca.balance -= amount
	ca.mu.Unlock()
	return nil
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)
	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))
	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
