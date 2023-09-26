package main

import (
	"fmt"
	"math"
)

const (
	loanAmount   = 10000
	interestRate = 0.05
	loanDuration = 48
)

func calculateMonthlyPayment(loanAmount float64, annualInterestRate float64, loanDuration int) float64 {
	monthlyInterestRate := annualInterestRate / 12
	monthlyPayment := (loanAmount * monthlyInterestRate) / (1 - math.Pow(1+monthlyInterestRate, float64(-loanDuration)))
	return monthlyPayment
}

func main() {
	monthlyPayment := calculateMonthlyPayment(loanAmount, interestRate, loanDuration)
	totalPayment := monthlyPayment * float64(loanDuration)

	fmt.Printf("Сумма кредита: $%.2f\n", loanAmount)
	fmt.Printf("Годовая процентная ставка: %.2f%%\n", interestRate*100)
	fmt.Printf("Срок кредита: %d месяцев\n", loanDuration)
	fmt.Printf("Ежемесячный платеж: $%.2f\n", monthlyPayment)
	fmt.Printf("Общая выплата: $%.2f\n", totalPayment)

}
