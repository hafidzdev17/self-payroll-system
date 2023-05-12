package model

import (
	"context"
	"time"
)

const (
	TransactionTypeDebit   = "debit"
	TransactionsTypeCredit = "credit"
)

type (
	Transaction struct {
		ID        int       `json:"id"`
		Amount    int       `json:"amount"`
		Note      string    `json:"note"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	TransactionRepository interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transaction, error)
	}

	TransactionUsecase interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transaction, int, error)
	}
)
