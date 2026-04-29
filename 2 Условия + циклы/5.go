package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Printf("Введите число 1: ")
	fmt.Scanln(&a)
	fmt.Printf("Введите число 2: ")
	fmt.Scanln(&b)
	fmt.Printf("Введите число 3: ")
	fmt.Scanln(&c)

	fmt.Printf("Наибольше число из (%d + %d = %d) -> %d\n", a, b, c, max(a, b, c))
}

func max(numbers ...int) int {
	var max = numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	return max
}
