package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "MTUCI"

	fmt.Printf("Длина символов: %d,\nСтрока включает в себя MTU? %t,\nИзменение регистра: %s", len(str), strings.Contains(str, "MTU"), strings.ToLower(str))
}
