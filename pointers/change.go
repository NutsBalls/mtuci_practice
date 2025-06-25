package main

import "fmt"

func change(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	first := 13
	second := 37

	fmt.Printf("Первое число равно %d, второе число равно %d\n", first, second)
	change(&first, &second)
	fmt.Printf("Первое число равно %d, второе число равно %d\n", first, second)
}
