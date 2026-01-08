package idempotency

import (
	"errors"
	"sync"
	"time"
)

var ErrDuplicateRequest = errors.New("duplicate request")

type MemoryStore struct {
	mu    sync.Mutex
	store map[string]Result
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]Result),
	}
}

func (s *MemoryStore) Get(key string) (Result, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	res, ok := s.store[key]
	return res, ok
}

func (s *MemoryStore) Set(key string, result Result) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[key] = result
}

func (s *MemoryStore) Reserve(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.store[key]; exists {
		return ErrDuplicateRequest
	}

	s.store[key] = Result{
		CreatedAt: time.Now(),
	}
	return nil
}
