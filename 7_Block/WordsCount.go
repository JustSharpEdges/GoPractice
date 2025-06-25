package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "АБА ОБА БЭМА БУМА АБА АБА БЭМА"

	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	fmt.Println("Результат подсчета:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
