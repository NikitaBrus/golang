package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	input := getUserIput()

	num, err := stringToNumber(input)

	if err == nil {
		fmt.Printf("Вы ввели число: %v\n", num)
	} else {
		fmt.Println("Это не число: ", err)
	}
}

func getUserIput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func stringToNumber(input string) (float64, error) {

	input = strings.Trim(input, "\" ' ")

	num, err := strconv.ParseFloat(input, 64)
	return num, err
}
