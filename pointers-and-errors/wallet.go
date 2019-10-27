package wallet

import "fmt"

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit method for Wallet struct
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance method for Wallet struct
func (w *Wallet) Balance() int {
	return w.balance
}
