package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	group := make(map[bool][]int)

	for _, number := range numbers {
		isEven := number%2 == 0
		group[isEven] = append(group[isEven], number)
	}

	fmt.Println(group[true])
	fmt.Println(group[false])
}
