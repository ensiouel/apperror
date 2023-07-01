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
	Code    codes.Code `json:"codes"`
	Message string     `json:"message"`
	Err     error      `json:"-"`
}

func New(code codes.Code) Error {
	return Error{
		Code:    code,
		Message: code.Message(),
	}
}

func (error Error) Error() string {
	if error.Err != nil {
		return error.Err.Error()
	}

	return error.Message
}

func (error Error) WithMessage(message string) Error {
	error.Message = message

	return error
}

func (error Error) WithError(err error) Error {
	error.Err = err

	return error
}

func (error Error) WithErrorf(format string, a ...interface{}) Error {
	error.Err = fmt.Errorf(format, a...)

	return error
}

func (error Error) Is(target error) bool {
	err, ok := target.(*Error)
	if !ok {
		return false
	}

	return error.Code == err.Code
}
