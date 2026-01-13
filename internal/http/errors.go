package http

import (
	"errors"
	"net/http"

	"github.com/dsrio/payment-processing-service/internal/domain"
	"github.com/dsrio/payment-processing-service/internal/idempotency"
)

func mapError(err error) int {
	switch {
	case errors.Is(err, domain.ErrInsufficientFunds):
		return http.StatusBadRequest
	case errors.Is(err, idempotency.ErrDuplicateRequest):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
