package main

import "fmt"

func main() {
	var slc []string

	slc = append(slc, "first", "second", "third", "fourth", "fifth")

	for _, value := range slc {
		fmt.Printf("Строка: %s\n", value)
	}
}
