package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Вот дом который построил Джек А это пшеница которая в темном чулане хранится В доме который построил Джек А это веселая птица-синица которая часто ворует пшеницу которая в темном чулане хранится В доме который построил Джек"

	arr := strings.Split(text, " ")

	countWords := make(map[string]int)

	for _, word := range arr {
		lword := strings.ToLower(word)
		countWords[lword]++
	}

	fmt.Println("Количество раз встречаемых в тексте слов:")
	for key, value := range countWords {
		fmt.Printf("%s: %d\n", key, value)
	}
}
