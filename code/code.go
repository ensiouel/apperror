package code

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

func (code Code) Message() string {
	return codeToMessage[code]
}
