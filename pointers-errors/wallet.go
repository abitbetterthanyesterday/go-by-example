package pointers

import (
	"errors"
	"fmt"
)

		var ErrorInsufficientFunds =  errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string

}
type Shitcoin int
func (s Shitcoin) String() string {
	return fmt.Sprintf("%d Shitcoin", s)
}

type Wallet struct {
	balance Shitcoin
}

func (w *Wallet) Deposit(amount Shitcoin){
	w.balance += amount
}

func (w Wallet) Balance() Shitcoin {
	return w.balance
}

func (w* Wallet) Withdraw(amount Shitcoin) error {
	
	if amount > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}