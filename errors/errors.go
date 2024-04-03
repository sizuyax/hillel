package errors

import "fmt"

var ErrEmptyBooks = Error{
	Code:    "INCORRECT_REQUEST",
	Message: "books cannot be empty!",
}

var ErrUnmarshalFail = Error{
	Code:    "SERVER_ERROR",
	Message: "failed unmarshal request",
}

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}
