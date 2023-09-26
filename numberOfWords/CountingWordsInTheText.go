package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите текст: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := scanner.Text()

	wordCount := countWords(inputText)

	fmt.Printf("Кол-ло свло в тексте: %d\n", wordCount)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
