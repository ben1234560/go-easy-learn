// Package errors provides a way to return detailed information
// for an RPC request error. The error is normally JSON encoded.
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error implements the error interface.
type Error struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func New(id, detail string, code int32) error {
	return &Error{
		Id:     id,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(id, format string, a ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   400,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(400),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, a ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   401,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(401),
	}
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, a ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   403,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(403),
	}
}

// NotFound generates a 404 error.
func NotFound(id, format string, a ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   404,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(404),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	return &Error{
		Id:     id,
		Code:   500,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(500),
	}
}
