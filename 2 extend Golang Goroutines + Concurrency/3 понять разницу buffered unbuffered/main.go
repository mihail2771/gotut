package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range ch {
			fmt.Printf("Задача %d выполнена\n", val+1)
		}
	}()

	for i := range 10 {
		ch <- i
		time.Sleep(1 + time.Second)
	}
	close(ch)

	wg.Wait()
}
