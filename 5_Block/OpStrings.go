package main

import (
	"fmt"
	"strings"
)

var Str string
var SubStr string

func main() {
	Str = "Наша строка номер один"
	SubStr = "стр"
	fmt.Println("Длина строки:", len(Str))
	fmt.Println("Str содержит SubStr?:", strings.Contains(Str, SubStr))
	fmt.Println("Вся строка в верхнем регистре:", strings.ToUpper(Str))
	fmt.Println("Вся строка в нижнем регистре:", strings.ToLower(Str))
	SplitedStr := strings.Split(Str, " ")
	fmt.Println("Слова в строке:", SplitedStr)
	for i, word := range SplitedStr {
		fmt.Println("Индекс слова:", i, ";", "Слово:", word)
	}
}
