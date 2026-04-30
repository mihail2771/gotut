package main

import "fmt"

func main() {
	var slice = []int{1, 4, 3, 2}
	var slice2 = []int{1, 4, 3, 2}
	fmt.Println(slice)
	fmt.Println(slice2)

	var result = append(slice, slice2...)
	fmt.Println(result)

}
