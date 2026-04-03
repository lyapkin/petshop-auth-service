package domain

import (
	"fmt"
)

type ErrorCode string

const (
	ErrNotFound         ErrorCode = "NOT_FOUND"
	ErrDuplicate        ErrorCode = "DUPLICATE"
	ErrNotValid         ErrorCode = "NOT_VALID"
	ErrNotAuthenticated ErrorCode = "NOT_AUTHENTICATED"

	ErrIDExists ErrorCode = "ID_EXISTS"

	ErrInternal ErrorCode = "INTERNAL"
)

const InternalErrorMessage = "Internal server error"

type AppError struct {
	Code     ErrorCode
	Message  string
	Internal error
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%v; %v", e.Message, e.Internal)
	}
	return e.Message
}

func NewInternalErr(err error) error {
	return &AppError{
		Code:     ErrInternal,
		Message:  InternalErrorMessage,
		Internal: err,
	}
}
