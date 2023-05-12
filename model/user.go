package model

import (
	"context"
	"self-payrol/request"
	"time"
)

type (
	User struct {
		ID         int       `json:"id"`
		SecretID   string    `json:"secret_id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Phone      string    `json:"phone"`
		Address    string    `json:"address"`
		PositionID int       `json:"position_id"`
		Position   *Position `json:"position"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	UserRepository interface {
		Create(ctx context.Context, user *User) (*User, error)
		UpdateByID(ctx context.Context, id int, user *User) (*User, error)
		FindByID(ctx context.Context, id int) (*User, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*User, error)
	}

	UserUsecase interface {
		GetByID(ctx context.Context, id int) (*User, error)
		FetchUser(ctx context.Context, limit, offset int) ([]*User, error)
		DestroyUser(ctx context.Context, id int) error
		EditUser(ctx context.Context, id int, req *request.UserRequest) (*User, error)
		StoreUser(ctx context.Context, req *request.UserRequest) (*User, error)
		WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error
	}
)
