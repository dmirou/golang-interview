package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v44/github"
)

func main() {
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	fmt.Println("err:", err)
	fmt.Printf("orgs %v\n", orgs)
}
