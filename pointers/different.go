package main

import "fmt"

func changeValue(n int) {
	n = 56
	fmt.Println("Изменено значение в функции:", n)
}

func changePointer(n *int) {
	*n = 78
	fmt.Println("Изменено значение при помощи указателя в функции:", *n)
}

func main() {
	number := 67
	fmt.Println("Исходное значене числа:", number)
	changeValue(number)
	fmt.Println("Значение изменилось только внутри функции:", number)
	changePointer(&number)
	fmt.Println("Значение изменилось:", number)
}
