package main

import (
	"fmt"
	"math"
)

var Pi = math.Pi
var E = math.E

func main() {
	fmt.Println("Значение Pi: ", Pi)
	fmt.Println("Значение E: ", E)

	const radius = 4.0
	circleArea := Pi * radius * radius
	fmt.Printf("Площадь круга радиуса %v равна: %.2f\n", radius, circleArea)

	Log := math.Log(E)
	fmt.Printf("Натуральный логарифм от E равен %v\n", Log)

}
