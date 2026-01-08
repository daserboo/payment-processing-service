package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpTransport "github.com/daserboo/payment-processing-service/internal/http"
	"github.com/daserboo/payment-processing-service/internal/idempotency"
	"github.com/daserboo/payment-processing-service/internal/locking"
	"github.com/daserboo/payment-processing-service/internal/repository/memory"
	"github.com/daserboo/payment-processing-service/internal/service"
)

func main() {
	// --- dependencies ---
	ledger := memory.NewLedgerRepository()
	locker := locking.NewAccountLocker()
	idem := idempotency.NewMemoryStore()

	paymentService := service.NewPaymentService(ledger, locker, idem)
	handler := httpTransport.NewPaymentHandler(paymentService)
	router := httpTransport.NewRouter(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// --- graceful shutdown ---
	go func() {
		log.Println("payments-api listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server exited cleanly")
}
