package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message,omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return newGenericError(message, http.StatusNotFound)
}

func NewUnexpectedError(message string) *AppError {
	return newGenericError(message, http.StatusInternalServerError)
}

func NewBadRequestError(message string) *AppError {
	return newGenericError(message, http.StatusBadRequest)
}

func NewValidationError(message string) *AppError {
	return newGenericError(message, http.StatusUnprocessableEntity)
}

func NewAuthorizationError(message string) *AppError {
	return newGenericError(message, http.StatusUnauthorized)
}

func newGenericError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}
