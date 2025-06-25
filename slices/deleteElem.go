package main

import (
	"fmt"
)

func deleteElem(slc []string, index int) []string {
	copy(slc[index:], slc[index+1:])
	slc[len(slc)-1] = ""
	slc = slc[:len(slc)-1]

	return slc
}

func main() {

	var slc []string

	slc = append(slc, "1", "2", "3", "4", "5")
	fmt.Printf("Срез до удаления элемента: %s\n", slc)
	index := 1

	slc = deleteElem(slc, index)

	fmt.Printf("Срез после удаления элемента по индексу %d: %s", index, slc)
}
