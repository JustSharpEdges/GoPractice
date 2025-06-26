package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	AvgGrade  float64
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "отличник"
	case s.AvgGrade >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.AvgGrade = newGrade
}

func main() {
	student := Student{
		Name:      "Иван Иванов",
		BirthYear: 2000,
		AvgGrade:  4.2,
	}

	fmt.Printf("%s: возраст %d лет\n", student.Name, student.GetAge())
	fmt.Printf("Статус: %s\n", student.GetStatus())

	student.UpdateGrade(4.7)
	fmt.Printf("Новый статус: %s\n", student.GetStatus())
}
