package main

import "fmt"

func main() {
	integer := 10
	float := 16.17
	var cmplx = complex(2, -2)
	var b byte = 65
	str := "hey"
	True := true

	fmt.Printf("Целое число: %d\n Число с плавающей точкой: %.2f\n Комплексный тип: %v\n Байт: %s\n Строки: %s\n Логические: %t", integer, float, cmplx, string(b), str, True)
}
