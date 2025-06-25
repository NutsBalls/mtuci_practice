package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Я сижу и делаю задание по практике"

	arr := strings.Split(str, " ")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
