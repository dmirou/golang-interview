package main

import (
	"errors"
	"fmt"

	"github.com/dmirou/learngo/error/handling/errortypes/internal/storage"
)

func TopLevelError() {
	err := storage.AddPerson("")
	// this package depends on the storage package because it uses his error
	var ae *storage.ActionError

	if errors.As(err, &ae) {
		fmt.Printf("action error: %s\n", ae)
		return
	}

	fmt.Printf("unknown error type: %s\n", err)
}

func ErrorInTheChain() {
	name := "Ivan"
	err := storage.AddPerson(name)

	// this package depends on the storage package because it uses his error
	var ae *storage.ActionError

	if errors.As(err, &ae) {
		fmt.Printf("action error: %s\n", ae)
		return
	}

	fmt.Printf("unknown error type: %s\n", err)
}

func ObfuscatedErrorInTheChain() {
	name := "Ivan"
	err := storage.DeletePerson(name)

	// this package depends on the storage package because it uses his error
	var ae *storage.ActionError

	if errors.As(err, &ae) {
		fmt.Printf("action error: %s\n", ae)
		return
	}

	fmt.Printf("unknown error type: %s\n", err)
}

func main() {
	TopLevelError()
	ErrorInTheChain()
	ObfuscatedErrorInTheChain()
}
