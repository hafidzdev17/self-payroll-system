package model

import (
	"context"
	"self-payrol/request"
	"time"
)

type (
	Company struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CompanyRepository interface {
		Get(ctx context.Context) (*Company, error)
		CreateOrUpdate(ctx context.Context, Company *Company) (*Company, error)
		AddBalance(ctx context.Context, balance int) (*Company, error)
		DebitBalance(ctx context.Context, amount int, note string) error
	}

	CompanyUsecase interface {
		GetCompanyInfo(ctx context.Context) (*Company, int, error)
		CreateOrUpdateCompany(ctx context.Context, req request.CompanyRequest) (*Company, int, error)
		TopupBalance(ctx context.Context, req request.TopupCompanyBalance) (*Company, int, error)
	}
)
