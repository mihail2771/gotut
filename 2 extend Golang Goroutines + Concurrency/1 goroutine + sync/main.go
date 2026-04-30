package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printVal(&wg)
	wg.Add(1)
	go printVal(&wg)
	wg.Add(1)
	go printVal(&wg)

	wg.Wait()

	wg.Add(1)
	go sum(&wg, 1, 3, 4, 3, 3, 4, 3, 2, 4, 3, 5, 7, 8, 8, 5, 4, 8, 9, 7, 6, 4, 4, 6, 7, 8, 5)
	wg.Add(1)
	go multiply(&wg, 1, 3, 4, 3, 3, 4, 3, 2, 4, 3, 5, 7, 8, 8, 5, 4, 8, 9, 7, 6, 4, 4, 6, 7, 8, 5)

	wg.Wait()
}

func sum(wg *sync.WaitGroup, numbers ...int) int {
	defer wg.Done()
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	defer fmt.Println(sum)
	return sum
}

func multiply(wg *sync.WaitGroup, numbers ...int) int {
	defer wg.Done()
	sum := 1
	for _, val := range numbers {
		sum *= val
	}
	defer fmt.Println(sum)
	return sum
}

func printVal(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 + time.Microsecond)
	}
}
