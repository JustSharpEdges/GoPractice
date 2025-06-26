package main

import (
	"fmt"
	"sort"
)

func FindElement(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func FilterSlice(slice []int, condition func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	target := 7
	index := FindElement(numbers, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}

	sorted := SortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	evenNumbers := FilterSlice(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Четные числа:", evenNumbers)

	largeNumbers := FilterSlice(numbers, func(x int) bool {
		return x > 5
	})
	fmt.Println("Числа больше 5:", largeNumbers)
}
