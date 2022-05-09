package storage

import (
	"errors"
	"fmt"

	appErr "github.com/dmirou/learngo/error/handling/opaque/internal/error"
)

var errConnection = errors.New("can not connect to the storage")
var errAlreadyAdded = errors.New("person already added")
var errAlreadyDeleted = errors.New("person already deleted")

type action string

const (
	// actionAddPerson is part of public API of the package.
	actionAddPerson action = "ADD_PERSON"
	// actionDeletePerson is part of public API of the package.
	actionDeletePerson action = "DELETE_PERSON"
)

// actionError is part of public API of the package.
type actionError struct {
	action     action
	code       appErr.Code
	args       map[string]interface{}
	wrappedErr error
}

func (ae *actionError) Error() string {
	return fmt.Sprintf(
		"action: %s, args: %v, code: %s, wrapped: %v",
		ae.action, ae.args, ae.code, ae.wrappedErr)
}

func (ae *actionError) Code() appErr.Code {
	return ae.code
}
