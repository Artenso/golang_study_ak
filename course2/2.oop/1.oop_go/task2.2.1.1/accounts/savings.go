package accounts

import "fmt"

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) Balance() float64 {
	return sa.balance
}

func (sa *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	sa.balance += amount
	return nil
}

func (sa *SavingsAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	if sa.balance < 500 {
		return fmt.Errorf("your current balance is lower than 500")
	}
	if sa.balance < amount {
		return fmt.Errorf("not enough mony")
	}
	sa.balance -= amount
	return nil
}
