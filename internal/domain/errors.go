package domain

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrAccountNotFound   = errors.New("account not found")
)
