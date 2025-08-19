package domain

import "errors"

var (
	ErrObjectTypeAlreadyExists = errors.New("object type already exists")
	ErrObjectTypeNotFound      = errors.New("object type not found")
)

type (
	ObjectTypeCreateRequest struct {
		Name string `json:"name" validate:"required,min=4"`
	}

	ObjectTypeUpdateRequest struct {
		Name *string `json:"name,omitempty" validate:"min=4"`
	}
)
