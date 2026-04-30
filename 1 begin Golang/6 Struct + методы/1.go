package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	user := User{Name: "Элизабет", Age: 26, IsActive: false}

	user.SayHi()
}

func (u *User) SayHi() {
	fmt.Printf("Привет, %s! Тебе уже %d лет, поздравляю!", u.Name, u.Age)
}
