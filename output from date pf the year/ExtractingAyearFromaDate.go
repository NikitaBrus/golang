package main

import (
	"fmt"
	"strings"
)

func main() {
	var date string
	fmt.Println("Введите дату в формате день.месяц.год ")
	fmt.Scan(&date)

	parts := strings.Split(date, ".")

	if len(parts) != 3 {
		fmt.Println("Неверный формат даты. Введите день.месяц.год")
		return
	}
	years := parts[2]
	fmt.Printf("Год: %s\n", years)
}
