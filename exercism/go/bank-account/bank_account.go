package account

import "sync"

type Account struct {
	monies int
	closed bool
	mutex  sync.Mutex
}

type account interface {
	Balance() (int, bool)
	Close() (int, bool)
	Deposit(int) (int, bool)
}

func Open(amount int) (a account) {
	if amount < 0 {
		return nil
	}

	a = &Account{amount, false, sync.Mutex{}}

	return
}

func (a *Account) Balance() (int, bool) {
	if a.closed {
		return 0, false
	}
	return a.monies, true
}

func (a *Account) Close() (int, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return 0, false
	} else {
		a.closed = true
	}

	return a.monies, true
}

func (a *Account) Deposit(amount int) (int, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return 0, false
	}

	if (a.monies + amount) < 0 {
		return 0, false
	}

	a.monies = a.monies + amount

	return a.monies, true
}
