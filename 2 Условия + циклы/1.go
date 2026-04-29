package main

import (
	"fmt"
)

func main() {
	type user struct {
		name  string
		email string
		age   int
	}
	e3 := user{
		name:  "na",
		email: "323",
		age:   13,
	}

	fmt.Println(e3.name)
	fmt.Println(e3.email)
	fmt.Println(e3.age)

}
