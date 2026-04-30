package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Летний вечер принес прохладу. Вечер был тихим, и этот вечер запомнился нам надолго. Мы шли через старый, старый парк. Парк встретил нас шумом листвы. Листва тихо шелестела, создавая уют. Это уютное место — наш любимый парк. Просто тихий парк, просто летний вечер, просто отдых."

	text = strings.ToLower(text)
	reg := regexp.MustCompile(`^[a-zа-я\s]`)

	text = reg.ReplaceAllString(text, "")

	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	for word, count := range counts {
		fmt.Printf("%s : %d\n", word, count)
	}
}
