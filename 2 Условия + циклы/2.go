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
		if n <= 0 {
			fmt.Println("Число не должно быть отрицательным или нулевым")
			count++
			continue
		} else {
			break
		}

	}

	fmt.Printf("Факториал числа %d равен %d\n", n, factorialN(n))
}

func factorialN(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}

	return sum
}
