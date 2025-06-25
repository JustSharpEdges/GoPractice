package main

import (
	"fmt"
	"strings"
)

var StudentMap = make(map[string]int)

func AddGrade(name string, grade int) {
	name = normalizeName(name)
	StudentMap[name] = grade
	fmt.Printf("Оценка %d добавлена для студента %s\n", grade, name)
}

func GetGrade(name string) (int, bool) {
	name = normalizeName(name)
	grade, exists := StudentMap[name]
	return grade, exists
}

func RemoveGrade(name string) {
	name = normalizeName(name)
	if _, exists := StudentMap[name]; exists {
		delete(StudentMap, name)
		fmt.Printf("Студент %s удален\n", name)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func PrintAllGrades() {
	fmt.Println("\nТекущие оценки:")
	for name, grade := range StudentMap {
		fmt.Printf("%s: %d\n", name, grade)
	}
}

func normalizeName(name string) string {
	return strings.Title(strings.ToLower(name))
}

func main() {
	// Добавление оценок
	AddGrade("Alice", 5)
	AddGrade("BOB", 5)
	AddGrade("alice", 5)

	// Поиск оценок
	if grade, exists := GetGrade("bOb"); exists {
		fmt.Printf("Оценка Bob: %d\n", grade)
	}

	// Удаление
	RemoveGrade("Bob")
	RemoveGrade("Charlie")

	// Вывод всех оценок
	PrintAllGrades()
}
