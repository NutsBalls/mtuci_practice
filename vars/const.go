package main

import "fmt"

const pi = 3.14
const e = 2.718

func main() {
	r := 10.00

	fmt.Printf("Площадь круга равна: %f \n", (r*r)*pi)

	fmt.Printf("Вычтем из радиуса число Эйлера: %f", r-e)
}
