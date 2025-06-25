package main

import (
	"fmt"
	"unsafe"
)

func main() {
	integer := 10
	float := 16.17
	var cmplx = complex(2, -2)
	var b byte = 65
	str := "hey"
	True := true

	fmt.Println("Размер целого числа в байтах:", unsafe.Sizeof(integer))
	fmt.Println("Размер числа с плавающей точкой в байтах:", unsafe.Sizeof(float))
	fmt.Println("Размер комплексного числа в байтах:", unsafe.Sizeof(cmplx))
	fmt.Println("Размер типа байт в байтах:", unsafe.Sizeof(b))
	fmt.Println("Размер строки в байтах:", unsafe.Sizeof(str))
	fmt.Println("Размер булевого значения в байтах:", unsafe.Sizeof(True))
}
