package storage

import "errors"

// ErrInvalidName is part of public API of the package.
var ErrInvalidName = errors.New("invalid name")

// AddPerson is a part of public API of the package.
func AddPerson(name string) error {
	if name == "" {
		return ErrInvalidName
	}

	// add person to some storage
	return nil
}
