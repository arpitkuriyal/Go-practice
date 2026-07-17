package bankaccount

import (
	"errors"
	"sync"
	"testing"
)

func TestAccountOperations(t *testing.T) {
	account, err := NewAccount(100)
	if err != nil {
		t.Fatal(err)
	}
	if err := account.Deposit(25); err != nil {
		t.Fatal(err)
	}
	if err := account.Withdraw(40); err != nil {
		t.Fatal(err)
	}
	if got, want := account.Balance(), 85.0; got != want {
		t.Fatalf("Balance() = %v, want %v", got, want)
	}
}

func TestAccountRejectsInvalidOperations(t *testing.T) {
	account, _ := NewAccount(10)

	if err := account.Deposit(0); !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("Deposit(0) error = %v, want ErrInvalidAmount", err)
	}
	if err := account.Withdraw(20); !errors.Is(err, ErrInsufficientFunds) {
		t.Fatalf("Withdraw(20) error = %v, want ErrInsufficientFunds", err)
	}
	if _, err := NewAccount(-1); !errors.Is(err, ErrInvalidInitialFund) {
		t.Fatalf("NewAccount(-1) error = %v, want ErrInvalidInitialFund", err)
	}
}

func TestAccountConcurrentDeposits(t *testing.T) {
	account, _ := NewAccount(0)
	const deposits = 100

	var wg sync.WaitGroup
	for range deposits {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := account.Deposit(1); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()

	if got := account.Balance(); got != deposits {
		t.Fatalf("Balance() = %v, want %v", got, deposits)
	}
}
