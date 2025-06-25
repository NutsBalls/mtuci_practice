package main

import (
	"fmt"
	"sort"
)

func findElement(slice []int, target int) (int, bool) {
	for index, value := range slice {
		if value == target {
			return index, true
		}
	}
	return -1, false
}

func sortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func filterForEvenNumbers(slice []int) []int {
	var result []int
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 4, 2, 6, 6, 5, 3, 8, 7, 9}

	fmt.Println("Срез:", numbers)
	if index, found := findElement(numbers, 8); found {
		fmt.Printf("Элемент найден на позиции %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	sorted := sortSlice(numbers)
	fmt.Println("Срез был отсортирован:", sorted)

	evenNumbers := filterForEvenNumbers(numbers)
	fmt.Println("Четные числа:", evenNumbers)
}
