package storage

import (
	"errors"
	"fmt"
)

var errConnection = errors.New("can not connect to the storage")
var errAlreadyAdded = errors.New("person already added")
var errAlreadyDeleted = errors.New("person already deleted")

type Action string

const (
	// ActionAddPerson is part of public API of the package.
	ActionAddPerson Action = "ADD_PERSON"
	// ActionDeletePerson is part of public API of the package.
	ActionDeletePerson Action = "DELETE_PERSON"
)

type Reason string

const (
	ReasonInvalidPersonName    Reason = "INVALID_PERSON_NAME"
	ReasonPersonAlreadyAdded   Reason = "PERSON_ALREADY_ADDED"
	ReasonPersonAlreadyDeleted Reason = "PERSON_ALREADY_DELETED"
	ReasonInternalError        Reason = "INTERNAL_ERROR"
)

// ActionError is part of public API of the package.
type ActionError struct {
	Action     Action
	Reason     Reason
	Args       map[string]interface{}
	WrappedErr error
}

func (ae *ActionError) Error() string {
	return fmt.Sprintf(
		"action: %s, args: %v, reason: %s, wrapped: %v",
		ae.Action, ae.Args, ae.Reason, ae.WrappedErr)
}
