package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Name string
	// a lot of fields
}

var (
	getConfigOnce = sync.OnceValues[Config, error](func() (Config, error) {
		defer fmt.Printf("OnceValues func called")
		return Config{
			Name: "first call",
		}, nil
	})
)

func getConfig() (*Config, error) {
	res, err := getConfigOnce()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			resCfg, err := getConfig()
			fmt.Println(resCfg, err)
		}()
	}

	wg.Wait()
}
