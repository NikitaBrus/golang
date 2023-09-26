package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	fmt.Printf("Текущая дата: %s\n", currentTime.Format("2023-09-22"))

	fmt.Printf("Текущий день недели: %s\n", currentTime.Weekday())
}
