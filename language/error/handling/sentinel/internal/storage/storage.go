package storage

import (
	"errors"
	"fmt"
)

// ErrInvalidName is part of public API of the package.
var ErrInvalidName = errors.New("invalid name")

// ErrConnection is part of public API of the package.
var ErrConnection = errors.New("can not connect to the storage")

// ErrAlreadyDeleted is part of public API of the package.
var ErrAlreadyDeleted = errors.New("person already deleted")

// AddPerson is a part of public API of the package.
func AddPerson(name string) error {
	if name == "" {
		return ErrInvalidName
	}

	err := addToStorage()
	if err != nil {
		return fmt.Errorf("add person: %w", err)
	}

	return nil
}

// DeletePerson is a part of public API of the package.
func DeletePerson(name string) error {
	if name == "" {
		return ErrInvalidName
	}

	err := deleteFromStorage()
	if err != nil {
		return fmt.Errorf("delete person: %v", err)
	}

	return nil
}

func addToStorage() error {
	return ErrConnection
}

func deleteFromStorage() error {
	return ErrAlreadyDeleted
}
