package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := `Ученик начал свой путь в мир программирования. 
	Он открыл редактор кода. Решал задачу на подсчет слов в тексте.
	И вот уже ученик завершил подсчет повторяющихся слов. Он улыбнулся, ощущая удовлетворение от выполненной задачи, 
	и готов был приступить к следующему этапу своего увлекательного путешествия в мир программирования.`

	sentanceWordCount := countWords(text)

	for word, count := range sentanceWordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
