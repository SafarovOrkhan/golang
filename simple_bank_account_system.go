package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	Owner   string
	Balance float64
	mu      sync.Mutex
}

func (self *BankAccount) Deposit(amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	self.mu.Lock()
	defer self.mu.Unlock()
	self.Balance += amount
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Deposited amount %v , Balance : %v\n", amount, self.Balance)
}

func (self *BankAccount) Withdraw(amount float64, wg *sync.WaitGroup) bool {
	defer wg.Done()
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.Balance < amount {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("Can't withdraw %v. Balance : %v\n", amount, self.Balance)
		return false
	} else {
		self.Balance -= amount
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("Withdraw complete %v . Balance: %v\n", amount, self.Balance)
		return true
	}
}

func (self *BankAccount) GetBalance(wg *sync.WaitGroup) float64 {
	defer wg.Done()
	self.mu.Lock()
	defer self.mu.Unlock()
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Balance : %v\n", self.Balance)
	return self.Balance
}

func main() {
	var wg sync.WaitGroup
	var acc1 BankAccount
	acc1 = BankAccount{Owner: "Lara", Balance: 0.0}

	for true {
		wg.Add(1)
		go acc1.Deposit((rand.Float64()*8)+7, &wg)
		wg.Add(1)
		go acc1.Withdraw((rand.Float64()*8)+7, &wg)
		wg.Add(1)
		go acc1.GetBalance(&wg)
		wg.Wait()
	}

}
