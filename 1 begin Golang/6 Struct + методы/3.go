package main

import (
	"fmt"
	"time"
)

type BankAccount struct {
	ID            int64
	OwnerName     string
	AccountNumber string
	Balance       float64
	Currency      string
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func main() {
	accounts := []BankAccount{
		{ID: 1, OwnerName: "Alice", AccountNumber: "ACC1001", Balance: 150000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, OwnerName: "Bob", AccountNumber: "ACC1002", Balance: 230000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, OwnerName: "Charlie", AccountNumber: "ACC1003", Balance: 50000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, OwnerName: "David", AccountNumber: "ACC1004", Balance: 780000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, OwnerName: "Eve", AccountNumber: "ACC1005", Balance: 120000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 6, OwnerName: "Frank", AccountNumber: "ACC1006", Balance: 99000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 7, OwnerName: "Grace", AccountNumber: "ACC1007", Balance: 450000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 8, OwnerName: "Heidi", AccountNumber: "ACC1008", Balance: 670000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 9, OwnerName: "Ivan", AccountNumber: "ACC1009", Balance: 300000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 10, OwnerName: "Judy", AccountNumber: "ACC1010", Balance: 15000, Currency: "USD", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, acc := range accounts {
		acc.getBalance()
		acc.deposit(100.0)
		acc.withdraw(300000.0)
	}

}

func (ba *BankAccount) getBalance() {
	fmt.Printf("У %s средств на счете %f\n", ba.OwnerName, ba.Balance)
}

func (ba *BankAccount) error(s string) {
	fmt.Printf("%s\n", s)
}

func (ba *BankAccount) deposit(val float64) {
	ba.Balance += val
	ba.getBalance()
}

func (ba *BankAccount) withdraw(val float64) {
	if ba.Balance-val < 0 {
		ba.error("Не хватает средств для снятия")
		return
	}

	ba.Balance -= val
	ba.getBalance()
}
