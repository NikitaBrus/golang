package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	movies := []string{
		"Звёздные войны",
		"Властелин колец",
		"Гарри Поттер",
		"Терминатор",
		"Интерстеллар",
		"Побег из Шоушенка",
		"Матрица",
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(movies))
	randomMovies := movies[randomIndex]

	fmt.Println("Случайный фильм: %s\n", randomMovies)
}
