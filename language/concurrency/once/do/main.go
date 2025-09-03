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
	cfg  Config
	once sync.Once
)

func getConfig() *Config {
	once.Do(func() {
		defer fmt.Printf("once.do called")
		cfg = Config{
			Name: "first call",
		}
	})

	once.Do(func() {
		defer fmt.Printf("once.do called")
		cfg = Config{
			Name: "second call",
		}
	})

	return &cfg
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
