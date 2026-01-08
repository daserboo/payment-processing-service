package http

import "net/http"

func NewRouter(handler *PaymentHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/payments", handler.CreatePayment)

	return mux
}
