package task12

import (
	"sync"
	"testing"
	"time"
)

func TestBankAccountBasic(t *testing.T) {
	account := NewBankAccount(100)

	if account.GetBalance() != 100 {
		t.Errorf("Expected initial balance 100, got %d", account.GetBalance())
	}

	// Тестируем депозит
	account.Deposit(50)
	if account.GetBalance() != 150 {
		t.Errorf("Expected balance 150 after deposit, got %d", account.GetBalance())
	}

	// Тестируем успешный вывод
	success := account.Withdraw(30)
	if !success {
		t.Error("Withdrawal should succeed")
	}
	if account.GetBalance() != 120 {
		t.Errorf("Expected balance 120 after withdrawal, got %d", account.GetBalance())
	}

	// Тестируем неуспешный вывод
	success = account.Withdraw(200)
	if success {
		t.Error("Withdrawal should fail - insufficient funds")
	}
	if account.GetBalance() != 120 {
		t.Errorf("Balance should remain 120 after failed withdrawal, got %d", account.GetBalance())
	}
}

func TestBankAccountConcurrentDeposits(t *testing.T) {
	account := NewBankAccount(0)
	const numGoroutines = 100
	const depositAmount = 1

	var wg sync.WaitGroup

	// Запускаем горутины, которые делают депозиты
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(depositAmount)
		}()
	}

	wg.Wait()

	expected := numGoroutines * depositAmount
	if account.GetBalance() != expected {
		t.Errorf("Expected balance %d after concurrent deposits, got %d", expected, account.GetBalance())
	}
}

func TestBankAccountConcurrentWithdrawals(t *testing.T) {
	account := NewBankAccount(1000)
	const numGoroutines = 100
	const withdrawalAmount = 1

	var wg sync.WaitGroup

	// Запускаем горутины, которые делают выводы
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Withdraw(withdrawalAmount)
		}()
	}

	wg.Wait()

	expected := 1000 - (numGoroutines * withdrawalAmount)
	if account.GetBalance() != expected {
		t.Errorf("Expected balance %d after concurrent withdrawals, got %d", expected, account.GetBalance())
	}
}

func TestBankAccountConcurrentMixedOperations(t *testing.T) {
	account := NewBankAccount(500)
	const numGoroutines = 50

	var wg sync.WaitGroup

	// Половина горутин делает депозиты, половина - выводы
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(deposit bool) {
			defer wg.Done()
			if deposit {
				account.Deposit(2)
			} else {
				account.Withdraw(1)
			}
		}(i%2 == 0)
	}

	wg.Wait()

	// Ожидаемое изменение: 25 депозитов по 2 + 25 выводов по 1 = +25
	expected := 500 + 25
	if account.GetBalance() != expected {
		t.Errorf("Expected balance %d after mixed operations, got %d", expected, account.GetBalance())
	}
}

func TestTransferBasic(t *testing.T) {
	from := NewBankAccount(100)
	to := NewBankAccount(50)

	success := Transfer(from, to, 30)
	if !success {
		t.Error("Transfer should succeed")
	}

	if from.GetBalance() != 70 {
		t.Errorf("From account should have balance 70, got %d", from.GetBalance())
	}
	if to.GetBalance() != 80 {
		t.Errorf("To account should have balance 80, got %d", to.GetBalance())
	}
}

func TestTransferInsufficientFunds(t *testing.T) {
	from := NewBankAccount(50)
	to := NewBankAccount(100)

	success := Transfer(from, to, 100)
	if success {
		t.Error("Transfer should fail - insufficient funds")
	}

	if from.GetBalance() != 50 {
		t.Errorf("From account balance should remain 50, got %d", from.GetBalance())
	}
	if to.GetBalance() != 100 {
		t.Errorf("To account balance should remain 100, got %d", to.GetBalance())
	}
}

func TestTransferConcurrent(t *testing.T) {
	account1 := NewBankAccount(1000)
	account2 := NewBankAccount(1000)
	const numTransfers = 50

	var wg sync.WaitGroup

	// Запускаем горутины, которые делают переводы в обе стороны
	for i := 0; i < numTransfers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Transfer(account1, account2, 10)
		}()
	}

	for i := 0; i < numTransfers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Transfer(account2, account1, 10)
		}()
	}

	wg.Wait()

	// Общий баланс должен остаться неизменным
	totalBalance := account1.GetBalance() + account2.GetBalance()
	expectedTotal := 2000
	if totalBalance != expectedTotal {
		t.Errorf("Total balance should remain %d, got %d", expectedTotal, totalBalance)
	}
}

func TestBankAccountNegativeBalance(t *testing.T) {
	account := NewBankAccount(100)

	// Пытаемся снять больше, чем есть на счету
	success := account.Withdraw(150)
	if success {
		t.Error("Withdrawal should fail - insufficient funds")
	}

	if account.GetBalance() != 100 {
		t.Errorf("Balance should remain 100 after failed withdrawal, got %d", account.GetBalance())
	}
}

func TestBankAccountZeroAmount(t *testing.T) {
	account := NewBankAccount(100)

	// Депозит нулевой суммы
	account.Deposit(0)
	if account.GetBalance() != 100 {
		t.Errorf("Balance should remain 100 after zero deposit, got %d", account.GetBalance())
	}

	// Вывод нулевой суммы
	success := account.Withdraw(0)
	if !success {
		t.Error("Zero withdrawal should succeed")
	}
	if account.GetBalance() != 100 {
		t.Errorf("Balance should remain 100 after zero withdrawal, got %d", account.GetBalance())
	}
}

func TestTransferRaceAndDeadlock(t *testing.T) {
	a := NewBankAccount(1000)
	b := NewBankAccount(1000)
	var wg sync.WaitGroup

	transfer := func(from, to *BankAccount, amount int) {
		defer wg.Done()
		ok := Transfer(from, to, amount)
		if !ok {
			t.Errorf("Transfer failed: from %v to %v amount %d", from, to, amount)
		}
	}

	// Запускаем несколько конкурентных переводов между двумя счетами (и туда, и обратно)
	numTransfers := 100
	for i := 0; i < numTransfers; i++ {
		wg.Add(2)
		go transfer(a, b, 1)
		go transfer(b, a, 1)
	}

	// Для проверки дедлока: тест должен завершаться за разумное время
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(5 * time.Second):
		t.Fatal("Test is potentially deadlocked or too slow")
	}

	total := a.GetBalance() + b.GetBalance()
	if total != 2000 {
		t.Errorf("Total balance mismatch: got %d, want %d", total, 2000)
	}
}

func TestTransferNotAtomic(t *testing.T) {
	a := NewBankAccount(100)
	b := NewBankAccount(100)
	var wg sync.WaitGroup

	// Две одновременных попытки снять весь баланс со счета `a` на счет `b`
	wg.Add(2)
	go func() {
		defer wg.Done()
		Transfer(a, b, 100)
	}()
	go func() {
		defer wg.Done()
		Transfer(a, b, 100)
	}()
	wg.Wait()

	balanceA := a.GetBalance()
	balanceB := b.GetBalance()

	if balanceA < 0 {
		t.Fatalf("Некорректный баланс: счёт ушел в минус (%d)! Ожидается неотрицательный остаток.", balanceA)
	}
	if balanceA+balanceB != 200 {
		t.Fatalf("Сумма всех денег должна быть 200, а не %d (a=%d, b=%d)", balanceA+balanceB, balanceA, balanceB)
	}
}

func TestTransferToSelf(t *testing.T) {
	acct := NewBankAccount(100)

	success := Transfer(acct, acct, 30)
	if !success {
		t.Fatalf("Transfer to self should succeed when enough funds")
	}

	want := 100 // баланс не должен меняться, фактически
	got := acct.GetBalance()
	if got != want {
		t.Errorf("Expected balance %d after self-transfer, got %d", want, got)
	}

	// Пробуем перевести больше, чем на счете
	failed := Transfer(acct, acct, 200)
	if failed {
		t.Fatalf("Transfer to self with insufficient funds should fail")
	}
}
