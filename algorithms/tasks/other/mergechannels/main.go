package main

import "sync"

// Написать функцию, которая получает неопределенное число каналов
// и соединяет их в один канал (func merge(cs ...<-chan int) <-chan int)
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	result := make(chan int)

	wg.Add(len(cs))
	for i := 0; i < len(cs); i++ {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
		}(cs[i])
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {

}
