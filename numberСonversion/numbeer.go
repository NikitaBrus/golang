package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b int
	var sum int
	fmt.Println("Введите числа: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum = a + b

	sumAsString := strconv.Itoa(sum)
	fmt.Printf("Результат: %s\n", sumAsString)
}
