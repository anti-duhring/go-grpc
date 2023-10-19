package errors

import (
	"fmt"
)

type Error struct {
	Code    int
	Message string
	err     error
}

func (e Error) Error() string {
	return e.Message
}

func New(code int, message string) Error {

	return Error{
		Code:    code,
		Message: message,
		err:     fmt.Errorf(message),
	}
}

func Is(err Error, code int) bool {
	return err.Code == code
}
