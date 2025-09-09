package task12

import "sync"

// Task 12: Race Condition Fix

// BankAccount представляет банковский счет
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// NewBankAccount создает новый банковский счет
func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{
		balance: initialBalance,
		mu:      sync.Mutex{},
	}
}

// Deposit добавляет деньги на счет
func (ba *BankAccount) Deposit(amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.balance += amount
}

// Withdraw снимает деньги со счета
func (ba *BankAccount) Withdraw(amount int) bool {
	ba.mu.Lock()
	defer ba.mu.Unlock()
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
	if from.Withdraw(amount) {
		to.Deposit(amount)
		return true
	}
	return false
}
