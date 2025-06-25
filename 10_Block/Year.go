package main

import "fmt"

func ismyYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if ismyYear(year) {
		fmt.Printf("%d - високосный год\n", year)
	} else {
		fmt.Printf("%d - не високосный год\n", year)
	}
}
