package main

import "fmt"

func main() {
	var n int
	var count int = 0
	const MAX_COUNT = 3

	for {
		if count >= MAX_COUNT {
			fmt.Printf("Истекли попытки ввода")
			return
		}

		fmt.Printf("Введите число: ")
		fmt.Scanln(&n)
		if n < 0 {
			fmt.Println("Число не должно быть отрицательным")
			count++
			continue
		} else {
			break
		}

	}

	fmt.Printf("Результат суммы от 0 до %d равно %d\n", n, sumZeroTo(n))
}

func sumZeroTo(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
