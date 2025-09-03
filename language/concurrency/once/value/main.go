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
	getConfigOnce = sync.OnceValue[Config](func() Config {
		defer fmt.Printf("OnceValue func called")
		return Config{
			Name: "first call",
		}
	})
)

func getConfig() *Config {
	res := getConfigOnce()
	return &res
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			resCfg := getConfig()
			fmt.Println(resCfg)
		}()
	}

	wg.Wait()
}
