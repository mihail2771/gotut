package main

import (
	"fmt"
	"math"
)

func main() {
	var number int

	fmt.Printf("Введите число: ")
	fmt.Scanln(&number)

	var result bool = isPrime(number)

	if result {
		fmt.Println("Число является натуральным")
	} else {
		fmt.Println("Число не является натуральным")
	}
}

func isPrime(number int) bool {

	if number < 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	max := int(math.Sqrt(float64(number)))

	for i := 5; i < max; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}
