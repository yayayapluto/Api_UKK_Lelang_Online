package domain

import "errors"

var (
	ErrOrganizerAlreadyExists = errors.New("object type already exists")
	ErrOrganizerNotFound      = errors.New("object type not found")
)

type (
	OrganizerCreateRequest struct {
		Name          string `json:"name" validate:"required,min=2"`
		Address       string `json:"address" validate:"required"`
		BankName      string `json:"bank_name" validate:"required"`
		AccountNumber string `json:"account_number" validate:"required"`
		AccountName   string `json:"account_name" validate:"required"`
	}

	OrganizerUpdateRequest struct {
		Name          *string `json:"name,omitempty" validate:"min=2"`
		Address       *string `json:"address,omitempty" validate:""`
		BankName      *string `json:"bank_name,omitempty" validate:""`
		AccountNumber *string `json:"account_number,omitempty" validate:""`
		AccountName   *string `json:"account_name,omitempty" validate:""`
	}
)
