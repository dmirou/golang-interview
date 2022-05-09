package storage

import (
	"errors"
)

// AddPerson is a part of public API of the package.
func AddPerson(name string) error {
	if name == "" {
		return &ActionError{
			Action: ActionAddPerson,
			Reason: ReasonInvalidPersonName,
			Args: map[string]interface{}{
				"name": name,
			},
		}
	}

	err := addToStorage()
	if err != nil {
		reason := ReasonInternalError
		if errors.Is(err, errAlreadyAdded) {
			reason = ReasonPersonAlreadyAdded
		}

		return &ActionError{
			Action: ActionAddPerson,
			Reason: reason,
			Args: map[string]interface{}{
				"name": name,
			},
			WrappedErr: err,
		}
	}

	return nil
}

// DeletePerson is a part of public API of the package.
func DeletePerson(name string) error {
	if name == "" {
		return &ActionError{
			Action: ActionDeletePerson,
			Reason: ReasonInvalidPersonName,
			Args: map[string]interface{}{
				"name": name,
			},
		}
	}

	err := deleteFromStorage()
	if err != nil {
		reason := ReasonInternalError
		if errors.Is(err, errAlreadyDeleted) {
			reason = ReasonPersonAlreadyDeleted
		}

		return &ActionError{
			Action: ActionDeletePerson,
			Reason: reason,
			Args: map[string]interface{}{
				"name": name,
			},
			WrappedErr: err,
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
