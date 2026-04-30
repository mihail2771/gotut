package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	test1()
	test2()
	test3()
}

func test1() {
	var wg sync.WaitGroup
	count := 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func test2() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func test3() {
	var wg sync.WaitGroup
	//var mu sync.Mutex
	var count int64
	count = 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
