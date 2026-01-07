package domain

import "time"

type EntryType string

const (
	EntryDebit  EntryType = "debit"
	EntryCredit EntryType = "credit"
	EntryRefund EntryType = "refund"
)

type LedgerEntry struct {
	ID        string
	AccountID string
	Amount    int64
	Currency  string
	Type      EntryType
	CreatedAt time.Time
}

// count balance
func CalculateBalance(entries []LedgerEntry) int64 {
	var balance int64
	for _, e := range entries {
		balance += e.Amount
	}
	return balance
}
