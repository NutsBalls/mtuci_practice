package main

import "fmt"

func main() {

	students := make(map[string]int)

	addStudent(students, "Влад", 4)
	addStudent(students, "Андрей", 4)
	addStudent(students, "Егор", 3)
	addStudent(students, "Максим", 5)
	fmt.Println("Список всех студентов")
	for key, value := range students {
		fmt.Printf("Студент с именем %s имеет оценку %d\n", key, value)
	}

	findStudent(students, "Влад")
	fmt.Println("")

	deleteStudent(students, "Влад")
	deleteStudent(students, "Егор")
	fmt.Println("Список всех студентов после удаления")
	for key, value := range students {
		fmt.Printf("Студент с именем %s имеет оценку %d\n", key, value)
	}
	findStudent(students, "Влад")

}

func addStudent(students map[string]int, name string, mark int) {
	students[name] = mark
}

func findStudent(students map[string]int, name string) {
	mark, ok := students[name]

	if !ok {
		fmt.Printf("Студент %s не найден\n", name)
	} else {
		fmt.Printf("Студент %s найден и имеет оценку %d\n", name, mark)
	}
}

func deleteStudent(students map[string]int, name string) {
	delete(students, name)
}
