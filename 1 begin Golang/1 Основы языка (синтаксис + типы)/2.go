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

	fmt.Printf("Сумма %d + %d = %d\n", a, b, sum(a, b))
	fmt.Printf("Среднее (%d + %d + %d) / 3 = %f\n", a, b, c, float32((a+b+c))/3)
}

func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func average3(numbers ...int) float32 {
	return float32(sum(numbers...)) / (float32(len(numbers)))
}
