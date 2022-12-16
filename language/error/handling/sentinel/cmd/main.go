package main

import (
	"errors"
	"fmt"

	"github.com/dmirou/learngo/error/handling/sentinel/internal/storage"
)

func TopLevelError() {
	err := storage.AddPerson("")
	// this package depends on the storage package because it uses his error
	if errors.Is(err, storage.ErrInvalidName) {
		fmt.Printf("can not add person to the storage: %v\n\n", err)
	}
}

func ErrorInTheChain() {
	name := "Ivan"
	err := storage.AddPerson(name)
	if err != nil {
		if errors.Is(err, storage.ErrConnection) {
			fmt.Println("ErrConnection found in the chain because %w format was used")
		}
		fmt.Printf("error in the chain: %v\n\n", err)
	}
}

func ObfuscatedErrorInTheChain() {
	name := "Ivan"
	err := storage.DeletePerson(name)
	if err != nil {
		if !errors.Is(err, storage.ErrAlreadyDeleted) {
			fmt.Println("ErrAlreadyDeleted not found in the chain because %v format was used")
		}
		fmt.Printf("obfuscated error: %v\n", err)
	}
}

func main() {
	TopLevelError()
	ErrorInTheChain()
	ObfuscatedErrorInTheChain()
}
