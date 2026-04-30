package main

import "fmt"

func main() {
	var array = []int{1, 4, 3, 2}
	fmt.Println(array)

	// Конец
	array = append(array, 3)
	fmt.Println(array)

	// Начало
	array = append([]int{1}, array...)
	fmt.Println(array)

	// несколько
	array = append(array, 3, 4, 3, 3)
	fmt.Println(array)

	// замена
	array[3] = 4
	fmt.Println(array)

	// удаление
	array = append(array[:3], array[4:]...)
	fmt.Println(array)

	// вставка в середину
	array = append(array, 0)
	fmt.Println("1 ", array)
	copy(array[4:], array[3:])
	fmt.Println("2 ", array)
	array[3] = 111
	fmt.Println("3 ", array)

	// очиста слайса clear(s) Go 1.21+ - обнуляет все значения, но сохраняет длину.
	// Удаление всех элементов: s = s[:0] (длина станет 0, но память под ним сохранится для будущего использования)
}
