package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Printf("Введите число: ")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return

	}

	var sum int = 0

	for i := 055; i <= num; i++ {
		sum += i
	}
	fmt.Printf("Сумма всех чисел от 1 до %d равна %f\n", num, sum)
}
