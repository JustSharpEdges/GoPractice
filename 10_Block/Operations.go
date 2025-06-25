package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)

	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)

	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("Деление на ноль невозможно")
	}

	if b != 0 {
		fmt.Printf("%.2f %% %.2f = %.2f\n", a, b, math.Mod(a, b))
	} else {
		fmt.Println("Остаток от деления на ноль невозможен")
	}

	fmt.Printf("%.2f ^ %.2f = %.2f\n", a, b, math.Pow(a, b))

	fmt.Printf("√(%.2f + %.2f) = %.2f\n", a, b, math.Sqrt(a+b))
}
