package accounts

import "fmt"

type CurrentAccount struct {
	balance float64
}

func (ca *CurrentAccount) Balance() float64 {
	return ca.balance
}

func (ca *CurrentAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	ca.balance += amount
	return nil
}

func (ca *CurrentAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("invalid amount: %v, want amount > 0", amount)
	}
	if ca.balance < amount {
		return fmt.Errorf("not enough mony")
	}
	ca.balance -= amount
	return nil
}
