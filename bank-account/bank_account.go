// Package account to simulate a bank account supporting opening/closing, withdrawals, and deposits of money
package account

import "sync"

// Constant declarations
const testVersion = 1

// Account type to repesent an Account with a balance and isOpen flag
type Account struct {
	balance int
	isOpen  bool
}

var cMutex sync.Mutex

// Open function returns an opened Account type, or nil if amt is < 0.
func Open(amt int) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{amt, true}
}

// Close function on Account type. Closes an account and returns a pay out amount and true/false result
func (a *Account) Close() (int, bool) {
	cMutex.Lock()
	defer cMutex.Unlock()

	if a.isOpen == false {
		return 0, false
	}

	a.isOpen = false
	return a.balance, true
}

// Deposit function on Account type. Takes an +/- int value and deposits/withdraws value to/from account
func (a *Account) Deposit(depAmt int) (int, bool) {
	cMutex.Lock()
	defer cMutex.Unlock()

	if a.isOpen && (a.balance+depAmt >= 0) {
		a.balance += depAmt
		return a.balance, true
	}

	return a.balance, false
}

// Balance function on Account type. Returns the current balance on account and true/false result
func (a *Account) Balance() (int, bool) {
	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}
