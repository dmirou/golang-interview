package storage

import (
	"errors"

	appErr "github.com/dmirou/learngo/error/handling/errortypes/internal/error"
)

// AddPerson is a part of public API of the package.
func AddPerson(name string) error {
	if name == "" {
		return &actionError{
			action: actionAddPerson,
			code:   appErr.CodeInvalidPersonName,
			args: map[string]interface{}{
				"name": name,
			},
		}
	}

	err := addToStorage()
	if err != nil {
		reason := appErr.CodeInternalError
		if errors.Is(err, errAlreadyAdded) {
			reason = appErr.CodePersonAlreadyAdded
		}

		return &actionError{
			action: actionAddPerson,
			code:   reason,
			args: map[string]interface{}{
				"name": name,
			},
			wrappedErr: err,
		}
	}

	return nil
}

// DeletePerson is a part of public API of the package.
func DeletePerson(name string) error {
	if name == "" {
		return &actionError{
			action: actionDeletePerson,
			code:   appErr.CodeInvalidPersonName,
			args: map[string]interface{}{
				"name": name,
			},
		}
	}

	err := deleteFromStorage()
	if err != nil {
		reason := appErr.CodeInternalError
		if errors.Is(err, errAlreadyDeleted) {
			reason = appErr.CodePersonAlreadyDeleted
		}

		return &actionError{
			action: actionDeletePerson,
			code:   reason,
			args: map[string]interface{}{
				"name": name,
			},
			wrappedErr: err,
		}
	}

	return nil
}

func addToStorage() error {
	return errConnection
}

func deleteFromStorage() error {
	return errAlreadyDeleted
}
