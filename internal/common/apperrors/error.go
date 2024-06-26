package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	Authorization ErrorType = "AUTHORIZATION" // 401 UnAuthorize
	BadRequest    ErrorType = "BAD_REQUEST"   // BadInput - 400
	Conflict      ErrorType = "CONFLICT"      // Already exists (eg, create account with existent email) - 409
	Internal      ErrorType = "INTERNAL"      // Server (500) and fallback apperrors - 500
	NoRows        ErrorType = "NO_ROWS"       // 404 Not Found
	Unprocessable ErrorType = "UNPROCESSABLE" // 422 Unprocessable Entity
)

type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case Authorization:
		return http.StatusUnauthorized
	case Unprocessable:
		return http.StatusUnprocessableEntity
	case NoRows:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

func NewNoRows() *Error {
	return &Error{
		Type:    NoRows,
		Message: "no rows found",
	}
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewUnprocessable() *Error {
	return &Error{
		Type:    Unprocessable,
		Message: "foreign key constraint violation",
	}
}
