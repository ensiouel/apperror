package apperror

import (
	"fmt"
	"github.com/ensiouel/apperror/code"
)

var (
	Unknown          = New(code.Unknown)
	BadRequest       = New(code.BadRequest)
	NotFound         = New(code.NotFound)
	AlreadyExists    = New(code.AlreadyExists)
	PermissionDenied = New(code.PermissionDenied)
	Unauthorized     = New(code.Unauthorized)
	Internal         = New(code.Internal)
)

type Error struct {
	Code    code.Code `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

func New(code code.Code) *Error {
	return &Error{
		Code:    code,
		Message: code.Message(),
	}
}

func (error *Error) Error() string {
	if error.Err != nil {
		return error.Err.Error()
	}

	return error.Message
}

func (error *Error) WithMessage(message string) *Error {
	error.Message = message

	return error
}

func (error *Error) WithError(err error) *Error {
	error.Err = err

	return error
}

func (error *Error) WithErrorf(format string, a ...interface{}) *Error {
	error.Err = fmt.Errorf(format, a...)

	return error
}

func (error *Error) Is(target error) bool {
	err, ok := target.(*Error)
	if !ok {
		return false
	}

	return error.Code == err.Code
}
