package task12

import (
	"sync"
)

// Task 12: Race Condition Fix

// SafeCounter - потокобезопасный счетчик
type SafeCounter struct {
	m sync.Mutex
	v int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v++
}

func (c *SafeCounter) Decrement() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v--
}

func (c *SafeCounter) GetValue() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v
}

var cnt = NewSafeCounter()

// BankAccount представляет банковский счет
type BankAccount struct {
	balance int
	mu      sync.Mutex
	id      int64
}

// NewBankAccount создает новый банковский счет
func NewBankAccount(initialBalance int) *BankAccount {
	cnt.Increment()
	return &BankAccount{
		balance: initialBalance,
		mu:      sync.Mutex{},
		id:      int64(cnt.GetValue()),
	}
}

// Deposit добавляет деньги на счет
func (ba *BankAccount) Deposit(amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.deposit(amount)
}

func (ba *BankAccount) deposit(amount int) {
	ba.balance += amount
}

// Withdraw снимает деньги со счета
func (ba *BankAccount) Withdraw(amount int) bool {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	return ba.withdraw(amount)
}

func (ba *BankAccount) withdraw(amount int) bool {
	// Проверьте, достаточно ли средств
	if ba.balance >= amount {
		ba.balance -= amount
		return true
	}
	return false
}

// GetBalance возвращает текущий баланс
func (ba *BankAccount) GetBalance() int {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	return ba.balance
}

// Transfer переводит деньги с одного счета на другой
func Transfer(from, to *BankAccount, amount int) bool {
	if from.id == to.id {
		from.mu.Lock()
		defer from.mu.Unlock()
		if from.withdraw(amount) {
			from.deposit(amount)
			return true
		}
		return false
	}

	if from.id < to.id {
		from.mu.Lock()
		defer from.mu.Unlock()
		to.mu.Lock()
		defer to.mu.Unlock()
	} else {
		to.mu.Lock()
		defer to.mu.Unlock()
		from.mu.Lock()
		defer from.mu.Unlock()
	}

	if from.withdraw(amount) {
		to.deposit(amount)
		return true
	}
	return false
}
