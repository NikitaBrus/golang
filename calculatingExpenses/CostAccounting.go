package main

import "fmt"

type Expense struct {
	Name   string
	Amount float64
}

func main() {

	expenses := []Expense{}

	for {
		var name string
		var amount float64
		fmt.Println("Введите название расхода (или 'exit' для выхода) ")
		fmt.Scan(&name)

		if name == "exit" {
			break
		}
		fmt.Println("Введите сумма расхода ")
		fmt.Scan(&amount)

		expense := Expense{Name: name, Amount: amount}
		expenses = append(expenses, expense)

		totalExpense := calculateTotalExpense(expenses)
		fmt.Printf("Общий расход: %.2f\n", totalExpense)
	}

}

func calculateTotalExpense(expenses []Expense) float64 {
	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}
