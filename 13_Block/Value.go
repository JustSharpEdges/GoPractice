package main

import "fmt"

func modifyByValue(val int) {
	val = val + 100
	fmt.Println("В процессе modifyByValue:", val)
}

func modifyByPointer(ptr *int) {
	*ptr = *ptr + 100
	fmt.Println("В процессе modifyByPointer:", *ptr)
}

func main() {
	original := 50

	fmt.Println("До modifyByValue:", original)
	modifyByValue(original)
	fmt.Println("После modifyByValue:", original)

	fmt.Println("\n До modifyByPointer:", original)
	modifyByPointer(&original)
	fmt.Println("После modifyByPointer:", original)
}
