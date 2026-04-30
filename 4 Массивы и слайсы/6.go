package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 4, 3, 2, 4, 3, 2, 4, 3, 2}

	counts := make(map[int]int)

	for _, n := range slice {
		counts[n]++ // Если ключа нет, Go вернет 0, и ++ сделает 1. Удобно!
	}

	for key, value := range counts {
		fmt.Printf("Элемент %v: %d раз\n", key, value)
	}

}
