package main

import "fmt"

func main() {

	integer := 3
	fmt.Printf("Таблица умножения для числа %d:\n", integer)
	for i := 1; i < 11; i++ {
		fmt.Printf("%d умноженное на %d = %d\n", integer, i, integer*i)
	}
}
