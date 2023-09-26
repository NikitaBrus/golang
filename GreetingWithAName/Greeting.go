package main

import "fmt"

func main() {
	var name string
	fmt.Println("Введите свое имя ")
	fmt.Scan(&name)
	fmt.Println("Привет " + name)
}
