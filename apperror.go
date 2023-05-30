package apperror

import "fmt"

type Code uint

const (
	Unknown Code = iota + 1
	BadRequest
	NotFound
	AlreadyExists
	PermissionDenied
	Unauthorized
	Internal
)

var codeToMessage = map[Code]string{
	Unknown:          "unknown",
	BadRequest:       "bad request",
	NotFound:         "not found",
	AlreadyExists:    "already exists",
	PermissionDenied: "permission denied",
	Unauthorized:     "unauthorized",
	Internal:         "internal",
}

type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func New(code Code) *Error {
	return &Error{
		Code:    code,
		Message: codeToMessage[code],
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
