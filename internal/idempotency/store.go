package idempotency

import "time"

type Result struct {
	Err       error
	CreatedAt time.Time
}

type Store interface {
	Get(key string) (Result, bool)
	Set(key string, result Result)
	Reserve(key string) error
}
