package codes

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
	Unknown:          "unknown error",
	BadRequest:       "bad request",
	NotFound:         "not found",
	AlreadyExists:    "already exists",
	PermissionDenied: "permission denied",
	Unauthorized:     "unauthorized",
	Internal:         "internal error",
}

func (code Code) Message() string {
	return codeToMessage[code]
}
