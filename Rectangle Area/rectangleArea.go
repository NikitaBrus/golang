package main

import "fmt"

func main() {
	var a float64
	var b float64
	var Area float64

	fmt.Println("Введите a ")
	fmt.Scan(&a)

	fmt.Println("Введите b ")
	fmt.Scan(&b)

	Area = a * b
	fmt.Println("Площадь прямоугольника ", Area)

}
