package domain

import (
	"fmt"
	"net/http"
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

// TODO: fix it, the layer should not know what type of communication is used
var HTTPCode = map[ErrorCode]int{
	ErrNotFound:         http.StatusNotFound,
	ErrDuplicate:        http.StatusBadRequest,
	ErrNotValid:         http.StatusBadRequest,
	ErrIDExists:         http.StatusInternalServerError,
	ErrInternal:         http.StatusInternalServerError,
	ErrNotAuthenticated: http.StatusUnauthorized,
}

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
