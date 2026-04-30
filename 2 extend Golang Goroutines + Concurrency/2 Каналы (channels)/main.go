package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	ch2 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 40 {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			ch2 <- val * 2
		}
		close(ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch2 {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
