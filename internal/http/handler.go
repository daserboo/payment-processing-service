package http

import (
	"encoding/json"
	"net/http"

	"github.com/daserboo/payment-processing-service/internal/service"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	idemKey := r.Header.Get("Idempotency-Key")
	if idemKey == "" {
		http.Error(w, "missing Idempotency-Key", http.StatusBadRequest)
		return
	}

	var req CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err := h.service.ProcessPayment(
		r.Context(),
		idemKey,
		req.AccountID,
		req.Amount,
		req.Currency,
	)

	if err != nil {
		http.Error(w, err.Error(), mapError(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreatePaymentResponse{
		Status: "ok",
	})
}
