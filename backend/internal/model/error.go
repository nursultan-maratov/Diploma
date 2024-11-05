package model

type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return e.message
}

func New(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Message() string {
	return e.message
}
