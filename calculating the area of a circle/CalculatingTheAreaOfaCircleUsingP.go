package main

import "fmt"

func main() {
	const p = 3.1415
	var radius float64
	var Area float64
	var S float64
	fmt.Println("Введите r ")
	fmt.Scan(&radius)
	Area = p * radius * radius
	fmt.Println("Площадь кргуа равна ", Area)
	S = 2 * p * radius
	fmt.Println("Окружность равна ", S)
}
