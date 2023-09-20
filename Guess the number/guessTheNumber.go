package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const minNumber = 1
	const maxNumber = 100
	secretNumber := rand.Intn(maxNumber-minNumber+1) + minNumber

	fmt.Println("Здарова, заебал, э. Вот игра тебе 'Угадай число!")
	fmt.Printf("Я загадал число от 1 до 100 \n", minNumber, maxNumber)

	var guess int
	attempts := 0

	for {
		fmt.Print("Введи свое число: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Ошибка ввода. Ты слово ввел, дурак! Надо число")
			continue
		}

		attempts++

		if guess < secretNumber {
			fmt.Println("Я загад число больше")
		} else if guess > secretNumber {
			fmt.Println("Я загадал меньше")
		} else {
			fmt.Println("Допер! Хорош!")
			break
		}

	}

}
