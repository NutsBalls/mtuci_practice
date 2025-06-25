package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

func infoAboutStudent(s *Student) {
	fmt.Printf("Стуеденту %s, %d лет. Он учится на %d курсе и имеет средний балл %.2f\n", s.Name, s.Age, s.Course, s.GPA)
}

func transfer(s *Student) {
	s.Age++
	s.Course++
}

func NewGpa(s *Student, newGPA float64) {
	s.GPA = newGPA
}
func main() {

	student := Student{Name: "Vlad", Age: 19, Course: 2, GPA: 4.89}

	infoAboutStudent(&student)

	transfer(&student)
	NewGpa(&student, 4.91)

	fmt.Println("Обновлена информация о студенте")
	infoAboutStudent(&student)

}
