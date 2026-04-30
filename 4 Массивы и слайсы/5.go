package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice = []int{1, 4, 3, 2, 4, 3, 2, 4, 3, 2}
	slices.Sort(slice)
	fmt.Println(slices.Compact(slice))

	slice = []int{1, 4, 3, 2, 4, 3, 2, 4, 3, 2}
	fmt.Println(Unique(slice))
}

func Unique(slice []int) []int {
	result := []int{}
	seen := make(map[int]bool)

	for _, val := range slice {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}

	return result
}
