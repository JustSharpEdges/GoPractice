package main

import "fmt"

func RemValue(slice []int, index int) []int {
	if index < 0 || len(slice) <= index {
		fmt.Println("Удаляю элемент", slice[index])
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func main() {
	MySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный срез", MySlice)
	MySlice = RemValue(MySlice, 2)
	fmt.Println("Срез без удаленного элемента", MySlice)
}
