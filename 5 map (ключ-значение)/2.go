package main

import "fmt"

func main() {

	phoneBook := make(map[string]string)

	add(phoneBook, "Олег", "1231231")
	add(phoneBook, "Инна", "2312313")
	add(phoneBook, "София", "4234243")
	add(phoneBook, "Евгений", "34235454")

	fmt.Println(phoneBook)

	fmt.Println(read(phoneBook, "София"))

	update(phoneBook, "София", "1111")

	fmt.Println(read(phoneBook, "София"))

	deletem(phoneBook, "София")

	fmt.Println(read(phoneBook, "София"))
}

func add(mapss map[string]string, key string, val string) {
	mapss[key] = val
}

func read(mapss map[string]string, key string) string {
	return mapss[key]
}

func update(mapss map[string]string, key string, val string) {
	add(mapss, key, val)
}

func deletem(mapss map[string]string, key string) {
	delete(mapss, key)
}
