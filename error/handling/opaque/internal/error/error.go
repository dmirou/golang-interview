package error

import (
	"errors"
)

type Code string

const (
	CodeInvalidPersonName    Code = "INVALID_PERSON_NAME"
	CodePersonAlreadyAdded   Code = "PERSON_ALREADY_ADDED"
	CodePersonAlreadyDeleted Code = "PERSON_ALREADY_DELETED"
	CodeInternalError        Code = "INTERNAL_ERROR"
)

type codeGetter interface {
	Code() Code
}

// GetCode returns error code. If error isn't implement codeGetter interface
// or error has internal error code it returns CodeInternalError.
func GetCode(err error) Code {
	var r codeGetter
	if errors.As(err, &r) {
		return r.Code()
	}

	return CodeInternalError
}
