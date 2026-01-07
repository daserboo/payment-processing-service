package domain

import "time"

type TransactionStatus string

const (
	StatusPending TransactionStatus = "pending"
	StatusSuccess TransactionStatus = "success"
	StatusFailed  TransactionStatus = "failed"
)

type Transaction struct {
	ID        string
	AccountID string
	Amount    int64
	Currency  string
	Status    TransactionStatus
	CreatedAt time.Time
}
