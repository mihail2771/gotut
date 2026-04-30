package main

import "fmt"

func main() {
	var c float32

	fmt.Printf("Введите количесто Градусов Цельсия: ")
	fmt.Scanln(&c)

	fmt.Printf("Результа перевода в Градусы Фаренгейта:  %f", toFahrenheit(c))
}

func toFahrenheit(c float32) float32 {
	return (9.0/5.0)*c + 32
}
