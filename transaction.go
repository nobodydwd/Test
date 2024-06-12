package main

import (
	"fmt"
	"time"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	TimeStamp string
	From      string
	To        string
	Amount    string
}

// NewTransaction creates a new transaction
func NewTransaction(from, to, amount string) Transaction {
	return Transaction{
		TimeStamp: time.Now().String(),
		From:      from,
		To:        to,
		Amount:    amount,
	}
}

// Transfer performs a transfer between accounts
func Transfer(from, to string, amount int) error {
	if balances[from] < amount {
		return fmt.Errorf("insufficient funds")
	}

	balances[from] -= amount
	balances[to] += amount

	fmt.Printf("Transferred %d from %s to %s\n", amount, from, to)
	return nil
}
