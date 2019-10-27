package wallet

import "fmt"

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
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance method for Wallet struct
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
