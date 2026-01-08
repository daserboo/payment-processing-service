# Payment Processing Service (Go)

A fintech-style payment processing service written in Go.  
This project demonstrates production-oriented backend architecture, idempotent payment processing, and ledger-based accounting.

The goal of the project is to showcase clean design, correctness, and backend fundamentals commonly required in fintech systems.

---

## Key Concepts Demonstrated

- Idempotent payment processing
- Ledger-based balance tracking
- Clean Architecture (layered design)
- Thread-safe in-memory storage
- Deterministic and testable business logic
- HTTP API without business logic leakage
- Concurrency-safe operations

---

## Architecture Overview

The service follows Clean / Hexagonal Architecture principles.
```
cmd/
  payments-api/

internal/
  domain/
  service/
  repository/
  repository/memory/
  idempotency/
  locking/
  http/
```
---

## API

### Create Payment

POST /payments

Headers:
- Content-Type: application/json
- Idempotency-Key: <unique-key>

Body:
```
{
  "account_id": "acc-1",
  "amount": 100,
  "currency": "USD"
}
```
Responses:
- 201 Created
- 200 OK (idempotent replay)
- 400 Bad Request
- 422 Unprocessable Entity

---

## Running Locally
```
go test ./...
go run ./cmd/payments-api
```
---

## Notes

This project focuses on backend architecture and fintech-related concerns such as idempotency and ledger-based accounting.
