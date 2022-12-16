package main

import (
	"fmt"

	appErr "github.com/dmirou/learngo/error/handling/opaque/internal/error"
	"github.com/dmirou/learngo/error/handling/opaque/internal/storage"
)

func TopLevelError() {
	err := storage.AddPerson("")
	errorHandler(err)
}

func ErrorInTheChain() {
	name := "Ivan"
	err := storage.AddPerson(name)
	errorHandler(err)
}

func ObfuscatedErrorInTheChain() {
	name := "Ivan"
	err := storage.DeletePerson(name)
	errorHandler(err)
}

func errorHandler(err error) {
	if code := appErr.GetCode(err); code != appErr.CodeInternalError {
		// handle different cases with specific reasons
		// e.g. bad request error, etc.
		fmt.Printf("HTTP code 400: {code:%s}\n", code)
		return
	}

	// handle internal error
	fmt.Printf("HTTP code 500: {code: %s}\n", "INTERNAL_ERROR")
}

func main() {
	TopLevelError()
	ErrorInTheChain()
	ObfuscatedErrorInTheChain()
}
