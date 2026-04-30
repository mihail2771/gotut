package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice = []int{1, 4, 3, 2}
	fmt.Println(slice)

	Reverse(slice)
	fmt.Println(slice)

	Reverse(slice)
	fmt.Println(slice)

	slices.Reverse(slice)
	fmt.Println(slice)
}

func Reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
