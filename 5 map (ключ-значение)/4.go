package main

import "fmt"

func main() {

	map1 := map[string]int{
		"A": 1,
		"Б": 2,
		"В": 3,
		"Г": 4,
	}

	map2 := make(map[int]string)

	for key, val := range map1 {
		map2[val] = key
	}

	fmt.Println(map1)
	fmt.Println(map2)
}
