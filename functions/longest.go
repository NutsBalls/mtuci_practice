package main

import "fmt"

func findLongest(str ...string) string {
	if len(str) == 0 {
		return ""
	}

	longest := str[0]

	for _, s := range str {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return fmt.Sprintf("Самая длинная строка: %s", longest)
}

func main() {
	fmt.Println(findLongest("Владислав", "Егор", "Максим", "Никита"))
	fmt.Println(findLongest("BMW", "Lada", "Mercedes"))
	fmt.Println(findLongest("МТУСИ", "МГУ", "МФТИ"))
	fmt.Println(findLongest(""))
}
