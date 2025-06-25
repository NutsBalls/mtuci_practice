package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name   string
	Year   int
	Course int
	GPA    float64
}

func infoAboutStudent(s *Student) {
	fmt.Printf("Студент под именем %s, родился в %d году. Он учится на %d курсе и имеет средний балл %.2f\n", s.Name, s.Year, s.Course, s.GPA)
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.Year
}

func (s Student) GetStatus() string {
	switch {
	case s.GPA >= 4.5:
		return "отличник"
	case s.GPA >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := Student{Name: "Vlad", Year: 2005, Course: 2, GPA: 4.89}

	infoAboutStudent(&student)

	fmt.Printf("Студенту %d лет и он %s", student.GetAge(), student.GetStatus())
}
