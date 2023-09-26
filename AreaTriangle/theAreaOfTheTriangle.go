package main

import "fmt"

func main() {
	var a float64
	var h float64
	fmt.Println("Введите a ")
	fmt.Scan(&a)
	fmt.Printf("Введите h ")
	fmt.Scan(&h)

	S := (a * h) / 2
	fmt.Printf("Площадь треугольник: %.2f\n", S)

}
