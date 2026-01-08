package http

type CreatePaymentRequest struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
}

type CreatePaymentResponse struct {
	Status string `json:"status"`
}
