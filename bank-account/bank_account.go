package account

import (
	"sync"
)

type Account struct {
	price	int64 
	mu		sync.Mutex

	closed  bool
} 

func Open(initialDeposit int64) *Account { 
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		price:	initialDeposit, 
		closed:	false,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock() 
	opend := a.closed 
	
	if opend {
		return 0, false
	}

	a.closed = true 

	return a.price, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.Lock() 
	defer a.mu.Unlock()

	opend := a.closed 
	money := a.price
	if opend {
		return 0, false
	}

	
	return money, true
}

// 存款
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock() 
	defer a.mu.Unlock()

	opend := a.closed 
	
	if opend ||  a.price + amount < 0{
		return 0, false
	}

	a.price += amount
	return a.price, true
}