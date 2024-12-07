package errors

import (
	"fmt"
	"log"
	"net/http"
)

//*************************************************************
//Error Types help with Rest vs Grpc error formats
type NotFoundError error
type UnexpectedError error
type ValidationError error
type AuthenticationError error
type AuthorizationError error
type ConflictError error
type FatalDatabaseErr error

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusConflict,
	}
}

func ToRestErr(err error) *AppError {

	toString := func(err error) string {
		return fmt.Sprintf("%s", err)
	}
	if err != nil {
		//err message defined in domain
		switch err.(type) {
		case NotFoundError:
			return NewNotFoundError(toString(err))
		case UnexpectedError:
			return NewUnexpectedError(toString(err))
		case ValidationError:
			return NewValidationError(toString(err))
		case AuthenticationError:
			return NewAuthenticationError(toString(err))
		case AuthorizationError:
			return NewAuthorizationError(toString(err))
		case ConflictError:
			return NewConflictError(toString(err))
		case FatalDatabaseErr:
			log.Fatalf("fatal database err: %v", err)
		default:
			//TODO: think about what to do in this case
			log.Fatalf("Unexpected error type: %v", err)
		}
	}
	return nil
}
