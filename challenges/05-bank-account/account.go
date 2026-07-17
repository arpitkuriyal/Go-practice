// Package bankaccount demonstrates validation, error handling, and safe shared state.
package bankaccount

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrInvalidAmount      = errors.New("amount must be greater than zero")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrInvalidInitialFund = errors.New("initial balance cannot be negative")
)

// Account is safe to use from multiple goroutines.
type Account struct {
	mu      sync.Mutex
	balance float64
}

func NewAccount(initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, ErrInvalidInitialFund
	}
	return &Account{balance: initialBalance}, nil
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.balance {
		return fmt.Errorf("withdraw %.2f: %w", amount, ErrInsufficientFunds)
	}
	a.balance -= amount
	return nil
}

func (a *Account) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
