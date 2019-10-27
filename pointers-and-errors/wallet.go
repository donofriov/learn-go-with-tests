package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin type from existing int
type Bitcoin int

// Stringer interface
type Stringer interface {
	String() string
}

// String method for type alias Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit method for Wallet struct
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance method for Wallet struct
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw method for Wallet struct
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
