package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	PositionRequest struct {
		Name   string `json:"name" validate:"required"`
		Salary int    `json:"salary" validate:"required"`
	}

	EditPositionRequest struct {
		PositionRequest
	}
)

func (req PositionRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Salary, validation.Required),
	)
}

func (req EditPositionRequest) Validate() error {
	return validation.ValidateStruct(
		&req.PositionRequest,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Salary, validation.Required),
	)
}
