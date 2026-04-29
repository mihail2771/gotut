package main

import "fmt"

func main() {
	var a int

	fmt.Printf("Введите число: ")
	fmt.Scanln(&a)

	if a == 0 {
		fmt.Println("Число равно нулю")
	} else if a%2 == 0 {
		fmt.Println("Число является четным")
	} else {
		fmt.Println("Число является нечетным")
	}
}
