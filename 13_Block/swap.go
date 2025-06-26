package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	fmt.Printf("До обмена: x = %d, y = %d\n", x, y)

	// Передаем указатели на переменные
	swap(&x, &y)

	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
}
