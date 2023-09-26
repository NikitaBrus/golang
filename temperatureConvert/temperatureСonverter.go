package main

import "fmt"

func main() {
	var celsius float64
	fmt.Println("Введите температуру в Цельсиях ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Printf("%.2f градусов Цельсия равны %.2f Градусам Фаренгейта \n", celsius, fahrenheit)
}
