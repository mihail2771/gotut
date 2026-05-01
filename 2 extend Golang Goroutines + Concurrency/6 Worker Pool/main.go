package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	//ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go worker("Bob", 5*time.Second, ch1, &wg)
	go worker("Kop", 10*time.Second, ch1, &wg)
	go worker("Dop", 15*time.Second, ch1, &wg)

	go workCreate("Soup", 1*time.Second, ch1)

	wg.Wait()
}

func worker(name string, t time.Duration, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Планировщик %s - Задача %d в работе\n", name, val)
		time.Sleep(t)
		fmt.Printf("Планировщик %s - Задача %d выполнена\n", name, val)
		time.Sleep(t / 2)
	}
}

func workCreate(name string, t time.Duration, ch chan int) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		time.Sleep(t)
		fmt.Printf("%s создал задачу %d\n", name, i)
		ch <- i
	}
}
