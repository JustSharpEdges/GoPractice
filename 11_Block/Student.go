package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     strings.TrimSpace(name),
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.AvgGrade = newGrade
	fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.AvgGrade)
}

func (s Student) HasScholarship() bool {
	return s.AvgGrade >= 4.5
}

func main() {
	student := NewStudent("Иван Иванов", 20, 2, 4.7)

	student.PrintInfo()
	fmt.Println()

	// Проверяем стипендию
	if student.HasScholarship() {
		fmt.Printf("%s получает стипендию\n", student.Name)
	} else {
		fmt.Printf("%s не получает стипендию\n", student.Name)
	}
	fmt.Println()

	// Повышаем курс
	student.Promote()

	// Обновляем средний балл
	student.UpdateGrade(4.3)

	// Проверяем стипендию снова
	if student.HasScholarship() {
		fmt.Printf("%s получает стипендию\n", student.Name)
	} else {
		fmt.Printf("%s не получает стипендию\n", student.Name)
	}
}
