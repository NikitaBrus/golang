package main

import (
	"fmt"
)

func main() {
	for {
		var a int
		var b int
		var c int

		var Area int

		fmt.Println("Введите числа: ")
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)

		Area = (a + b + c) / 3
		fmt.Println("Результат: ", Area)

	}
}
