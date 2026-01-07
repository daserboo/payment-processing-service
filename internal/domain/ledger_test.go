package domain

import "testing"

func TestCalculateBalance(t *testing.T) {
	entries := []LedgerEntry{
		{Amount: 1000},
		{Amount: -300},
		{Amount: -200},
	}

	balance := CalculateBalance(entries)

	if balance != 500 {
		t.Fatalf("expected balance 500, got %d", balance)
	}
}
