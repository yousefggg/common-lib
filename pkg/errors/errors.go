package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInternal = errors.New("internal server error")
	ErrNotFound = errors.New("resource not found")
	ErrInvalidArgs = errors.New("invalid arguments")
	ErrUnauthorized = errors.New("unauthorized")
)
type CustomError struct {
	Code string 
	Message string 
	Err error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
//Unwrap - распаковка ошибок 
func (e *CustomError) Unwrap() error {
	return e.Err
}

func NewErr(code, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}