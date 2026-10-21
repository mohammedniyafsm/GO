package main

import (
	"fmt"
	"sync"
)

// ============================================
// 1. WHAT IS A MUTEX?
// ============================================
//
// Mutex means MUTUAL EXCLUSION.
//
// A Mutex allows only one goroutine at a time
// to access a protected section of code.
//
// Why do we need it?
//
// Imagine two customers withdrawing money from
// the SAME bank account at the same time.
//
// Account balance = 1000
//
// Customer A wants to withdraw 700.
// Customer B wants to withdraw 700.
//
// Both customers must not spend the same money.
//
// Without a Mutex:
// Both goroutines might read 1000.
// Both might decide that 700 is available.
// Both might withdraw based on the same balance.
//
// This is a concurrency problem.
//
// With a Mutex:
// One goroutine gets the lock first.
// It checks the balance and withdraws 700.
// The other goroutine waits for the lock.
// When it gets the lock, it sees only 300.
// The second withdrawal is rejected.
//
// Remember:
//
// Goroutines = Execute tasks concurrently.
// Mutex     = Protect shared data.
// WaitGroup = Wait for goroutines to finish.
// Channels  = Communicate values between goroutines.
//
// ============================================
// 2. IMPORTANT MUTEX METHODS
// ============================================
//
// Go provides Mutex inside the sync package.
//
// var mutex sync.Mutex
//
// Lock():
// Acquires the lock.
//
// mutex.Lock()
//
// If another goroutine already holds the lock,
// this goroutine waits until the lock is available.
//
// Unlock():
// Releases the lock.
//
// mutex.Unlock()
//
// Another goroutine can acquire the lock afterward.
//
// IMPORTANT:
//
// Lock() does not stop every goroutine in the program.
// It only makes other goroutines wait when they try
// to acquire the SAME mutex.
//
// ============================================
// 3. REAL-WORLD EXAMPLE: BANK ACCOUNT
// ============================================

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

// ============================================
// 4. DEPOSIT MONEY
// ============================================
//
// Multiple goroutines might deposit money into
// the same account.
//
// We protect the balance using the Mutex.

func (account *BankAccount) Deposit(amount int) {
	account.mutex.Lock()

	// Only one goroutine holding this Mutex
	// can execute this protected operation at a time.
	account.balance += amount

	account.mutex.Unlock()
}

// ============================================
// 5. WITHDRAW MONEY SAFELY
// ============================================
//
// The entire operation must be protected:
//
// 1. Check the balance.
// 2. Decide whether withdrawal is possible.
// 3. Update the balance.
//
// Why lock the entire operation?
//
// If we protect only the final subtraction,
// two goroutines could both check the same old balance
// and both decide that enough money is available.
//
// The check and update together form one critical section.

func (account *BankAccount) Withdraw(amount int, customer string) {
	account.mutex.Lock()

	// The lock is acquired before checking the balance.

	fmt.Println(customer, "is checking the balance")

	if account.balance >= amount {
		account.balance -= amount

		fmt.Println(customer, "withdrawal successful:", amount)
		fmt.Println("Remaining balance:", account.balance)
	} else {
		fmt.Println(customer, "withdrawal failed: insufficient funds")
		fmt.Println("Available balance:", account.balance)
	}

	// Release the lock after checking and updating.
	account.mutex.Unlock()
}

// ============================================
// 6. MAIN FUNCTION
// ============================================

func main() {
	// Start with 1000 in the bank account.
	account := BankAccount{
		balance: 1000,
	}

	// WaitGroup tracks both customer goroutines.
	var wg sync.WaitGroup

	wg.Add(2)

	// Two customers attempt to withdraw 700.
	go func() {
		defer wg.Done()
		account.Withdraw(700, "Customer A")
	}()

	go func() {
		defer wg.Done()
		account.Withdraw(700, "Customer B")
	}()

	// Wait until both withdrawals finish.
	wg.Wait()

	// Read the final balance safely.
	account.mutex.Lock()
	fmt.Println("\nFinal account balance:", account.balance)
	account.mutex.Unlock()
}

// ============================================
// EXPECTED OUTPUT: ONE POSSIBLE ORDER
// ============================================
//
// Customer A is checking the balance
// Customer A withdrawal successful: 700
// Remaining balance: 300
// Customer B is checking the balance
// Customer B withdrawal failed: insufficient funds
// Available balance: 300
//
// Final account balance: 300
//
// Customer B could run first instead.
// Goroutine execution order is not guaranteed.
//
// But regardless of the order:
// - Only one withdrawal succeeds.
// - The other withdrawal is rejected.
// - The final balance is 300.
//
// ============================================
// END OF PROGRAM
// ============================================
