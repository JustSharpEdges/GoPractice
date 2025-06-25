package main

import "fmt"

var mySlice []string

func main() {
	mySlice = append(mySlice, "first")
	mySlice = append(mySlice, "second")
	mySlice = append(mySlice, "third")
	fmt.Println(mySlice)
	fmt.Println("-------------------")
	fmt.Println("Или так")
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
}
