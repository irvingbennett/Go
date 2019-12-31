// Package account manages bank accounts
package account

import (
	"strconv"
	"sync"
)

var accounts = []string{}
var lastID = 0

// Account holds the balance of an account
type Account struct {
	ID      string
	balance int64
	mux     sync.Mutex
}

// Open return an bank accout
func Open(start int64) *Account {
	if start < 0 {
		return nil
	}
	var a Account
	a.ID = strconv.Itoa(lastID + 1)
	accounts = append(accounts, a.ID)
	lastID++
	a.balance = start
	return &a
}

// Close returns balance and closes account
func (a *Account) Close() (payout int64, ok bool) {
	if a.balance < 0 {
		return 0, false
	}
	for x := 0; x < len(accounts); x++ {
		if accounts[x] == a.ID {
			accounts[x] = ""
			payout = a.balance
			a.balance = 0
			a = nil
			ok = true
			return
		}
	}
	return 0, false
}

// Balance return the balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	for x := 0; x < len(accounts); x++ {
		if accounts[x] == a.ID {
			balance = a.balance
			ok = true
			return
		}
	}
	return 0, false
}

// Deposit adds to the balance of account and returns newBalance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	if amount < 0 && amount*-1 > a.balance {
		a.mux.Unlock()
		return 0, false
	}
	for x := 0; x < len(accounts); x++ {
		if accounts[x] == a.ID {
			a.balance += amount
			newBalance = a.balance
			ok = true
			a.mux.Unlock()
			return
		}
	}
	a.mux.Unlock()
	return 0, false
}
