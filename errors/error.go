package errors

import "errors"

var (
	ErrDataExistedDetail string = "data has been existed"
	ErrDataNotFound             = errors.New("data not found")
)

type Error struct {
	Code   Code
	Detail string
}

func (e Error) Error() string {
	msg := e.Code.String()
	if len(e.Detail) > 0 {
		msg += ": " + e.Detail
	}
	return msg
}

func ErrorIs(err error, taget Code) bool {
	if e, ok := err.(Error); ok {
		return e.Code == taget
	}
	return false
}
