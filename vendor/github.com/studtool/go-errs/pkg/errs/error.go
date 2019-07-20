package errs

//go:generate easyjson

import (
	"fmt"

	"github.com/mailru/easyjson"
)

const (
	Internal         = 0
	BadFormat        = 1
	InvalidFormat    = 2
	Conflict         = 3
	NotFound         = 4
	NotAuthorized    = 5
	PermissionDenied = 6
	NotImplemented   = 7
)

//easyjson:json
type Error struct {
	Type int `json:"-"`

	Code    int    `json:"code"`
	Message string `json:"message"`

	//nolint:govet
	cause error `json:"-"`

	//nolint:govet
	json []byte `json:"-"`

	//nolint:govet
	string string `json:"-"`
}

const (
	InternalErrorCode    = Internal
	InternalErrorMessage = "internal error"
)

func Wrap(err error) *Error {
	return NewCustom(
		Internal, InternalErrorCode,
		InternalErrorMessage, err,
	)
}

func NewCustom(
	t int, code int, message string, cause error,
) *Error {
	e := &Error{
		Type:    t,
		Code:    code,
		Message: message,
		cause:   cause,
	}

	if t == Internal {
		e.makeString()
	} else {
		e.makeJSON()
	}

	return e
}

func NewInternal(code int, message string) *Error {
	return NewCustom(
		InternalErrorCode, code, message, nil,
	)
}

func NewBadFormat(code int, message string) *Error {
	return NewCustom(
		BadFormat, code, message, nil,
	)
}

func NewInvalidFormat(code int, message string) *Error {
	return NewCustom(
		InvalidFormat, code, message, nil,
	)
}

func NewConflict(code int, message string) *Error {
	return NewCustom(
		Conflict, code, message, nil,
	)
}

func NewNotFound(code int, message string) *Error {
	return NewCustom(
		NotFound, code, message, nil,
	)
}

func NewNotAuthorized(code int, message string) *Error {
	return NewCustom(
		NotAuthorized, code, message, nil,
	)
}

func NewPermissionDenied(code int, message string) *Error {
	return NewCustom(
		PermissionDenied, code, message, nil,
	)
}

func NewNotImplemented(code int, message string) *Error {
	return NewCustom(
		NotImplemented, code, message, nil,
	)
}

func (v *Error) JSON() []byte {
	return v.json
}

func (v *Error) Error() string {
	return v.string
}

func (v *Error) Cause() error {
	return v.cause
}

func (v *Error) makeJSON() {
	v.json, _ = easyjson.Marshal(v)
}

func (v *Error) makeString() {
	v.string = fmt.Sprintf(
		"error: [type = '%d'; code: = '%d'; message = '%s'",
		v.Type, v.Code, v.Message,
	)
	if v.cause != nil {
		v.string += fmt.Sprintf("; cause = '%s']", v.cause.Error())
	} else {
		v.string += "]"
	}
}
