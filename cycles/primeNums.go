package main

import "fmt"

func main() {

	var primeNums []int

	for i := 2; i < 101; i++ {

		prime := true
		for _, value := range primeNums {
			if value%value > i {
				break
			}

			if i%value == 0 {
				prime = false
				break
			}
		}

		if prime {
			primeNums = append(primeNums, i)
		}
	}

	fmt.Println(primeNums)
}
