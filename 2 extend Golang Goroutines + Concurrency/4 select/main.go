package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		i := 1
		for {
			ch1 <- fmt.Sprintf("Быстро обработал задачу %d", i)
			time.Sleep(7 * time.Second)
			i++
		}
	}()

	go func() {
		i := 1
		for {
			ch2 <- fmt.Sprintf("Медленно обработал задачу %d", i)
			time.Sleep(10 * time.Second)
			i++
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("Таймаут: нет данных 2 секунды")
		}
	}
}
