package wallet

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var NegativeDepositError = errors.New("no negative amount allowed")
var OverdrawError = errors.New("cannot withdraw an amount greater than actual balance")

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return NegativeDepositError
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return OverdrawError
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
