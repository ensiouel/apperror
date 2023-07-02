package apperror

import (
	"fmt"
	"github.com/ensiouel/apperror/codes"
)

var (
	Unknown          = New(codes.Unknown)
	BadRequest       = New(codes.BadRequest)
	NotFound         = New(codes.NotFound)
	AlreadyExists    = New(codes.AlreadyExists)
	PermissionDenied = New(codes.PermissionDenied)
	Unauthorized     = New(codes.Unauthorized)
	Internal         = New(codes.Internal)
)

type Error struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
	Err     error      `json:"-"`
}

func New(code codes.Code) *Error {
	return &Error{
		Code:    code,
		Message: code.Message(),
	}
}

func (error *Error) clone() *Error {
	c := *error
	return &c
}

func (error *Error) Error() string {
	if error.Err != nil {
		return error.Err.Error()
	}

	return error.Message
}

func (error *Error) WithMessage(message string) *Error {
	c := error.clone()
	c.Message = message
	return c
}

func (error *Error) WithMessagef(format string, a ...interface{}) *Error {
	c := error.clone()
	c.Message = fmt.Sprintf(format, a...)
	return c
}

func (error *Error) WithError(err error) *Error {
	c := error.clone()
	c.Err = err
	return c
}

func (error *Error) WithErrorf(format string, a ...interface{}) *Error {
	c := error.clone()
	c.Err = fmt.Errorf(format, a...)
	return c
}

func (error *Error) Is(target error) bool {
	err, ok := target.(*Error)
	if !ok {
		return false
	}

	return error.Code == err.Code
}
