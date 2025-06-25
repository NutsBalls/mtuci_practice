package main

import "fmt"

func main() {
	var a, b float64
	var operator string
	fmt.Scan(&a, &b, &operator)

	switch operator {

	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
		} else {
			fmt.Println(a / b)
		}

	}

}
