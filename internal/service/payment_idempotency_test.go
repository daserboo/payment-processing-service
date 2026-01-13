package service

import (
	"context"
	"testing"

	"github.com/dsrio/payment-processing-service/internal/domain"
	"github.com/dsrio/payment-processing-service/internal/idempotency"
	"github.com/dsrio/payment-processing-service/internal/locking"
	"github.com/dsrio/payment-processing-service/internal/repository/memory"
)

func TestPaymentService_Idempotency(t *testing.T) {
	ledger := memory.NewLedgerRepository()
	locker := locking.NewAccountLocker()
	idem := idempotency.NewMemoryStore()

	service := NewPaymentService(ledger, locker, idem)

	_ = ledger.Append(domain.LedgerEntry{
		AccountID: "acc-1",
		Amount:    1000,
		Type:      domain.EntryCredit,
	})

	err1 := service.ProcessPayment(
		context.Background(),
		"idem-123",
		"acc-1",
		500,
		"USD",
	)

	err2 := service.ProcessPayment(
		context.Background(),
		"idem-123",
		"acc-1",
		500,
		"USD",
	)

	if err1 != nil {
		t.Fatal(err1)
	}
	if err2 != nil {
		t.Fatalf("second call must be idempotent")
	}
}
