package model

import (
	"context"
	"self-payrol/request"
	"time"
)

type (
	Position struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Salary    int       `json:"salary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	PositionRepository interface {
		Create(ctx context.Context, Position *Position) (*Position, error)
		UpdateByID(ctx context.Context, id int, Position *Position) (*Position, error)
		FindByID(ctx context.Context, id int) (*Position, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*Position, error)
	}

	PositionUsecase interface {
		GetByID(ctx context.Context, id int) (*Position, error)
		FetchPosition(ctx context.Context, limit, offset int) ([]*Position, error)
		DestroyPosition(ctx context.Context, id int) error
		EditPosition(ctx context.Context, id int, req *request.PositionRequest) (*Position, error)
		StorePosition(ctx context.Context, req *request.PositionRequest) (*Position, error)
	}
)
