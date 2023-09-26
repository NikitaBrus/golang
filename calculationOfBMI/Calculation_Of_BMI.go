package main

import "fmt"

func main() {
	var weight float64
	var height float64

	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&weight)

	fmt.Print("Введите ваш рост: ")
	fmt.Scan(&height)

	bmi := weight / (height * height)

	fmt.Printf("Ваш индекс ИМТ %.2f\n ", bmi)
	fmt.Println("Интерпретация ИМТ:")
	fmt.Println("Менее 18.5: Недостаточная масса тела")
	fmt.Println("18.5 - 24.9: Нормальная масса тела")
	fmt.Println("25.0 - 29.9: Избыточная масса тела (предожирение)")
	fmt.Println("30.0 и более: Ожирение")

}
