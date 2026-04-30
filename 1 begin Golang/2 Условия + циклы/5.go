package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	fibonacciTo(n)
}

func fibonacciTo(n int) {
	a, b := 0, 1

	fmt.Printf("Числа Фибоначчи до %d: [", n)

	first := true
	for a < n {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(a)

		a, b = b, a+b
		first = false
	}

	fmt.Println("]")
}
