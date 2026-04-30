package main

import "fmt"

func main() {
	var array = []int{} //{1,4,3,2,3,31,5,4,3,23,54,23,4,4,3,22,2,3,11,4}
	var array2 = []int{1, 4, 3, 2, 3, 31, 5, 4, 3, 23, 54, 23, 4, 4, 3, 22, 2, 3, 11, 4}

	var min, max int = minmax(array)
	fmt.Printf("max: %d \n", max)
	fmt.Printf("min: %d \n", min)

	min, max = minmax(array2)
	fmt.Printf("max: %d \n", max)
	fmt.Printf("min: %d \n", min)
}

func minmax(array []int) (int, int) {
	var min, max int
	if len(array) > 0 {
		min, max = array[0], array[0]
	} else {
		return 0, 0
	}

	for _, number := range array {
		if min > number {
			min = number
		}
		if max < number {
			max = number
		}
	}
	return min, max

}
