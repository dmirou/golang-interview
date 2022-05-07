package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dmirou/learngo/error/handling/sentinel/internal/storage"
)

func main() {
	err := storage.AddPerson("")
	// this package depends on the storage package because it uses his error
	if errors.Is(err, storage.ErrInvalidName) {
		fmt.Println("can not add person to the storage")
		os.Exit(1)
	}

	fmt.Println("saved successfully")
}
