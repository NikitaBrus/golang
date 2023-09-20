package main

import "fmt"

const (
	EnglishGreeting = "Hello!"
	SpanishGreeting = "Hola"
	FrenchGreeting  = "Salut"
	RussianGreeting = "Привет"
	GermanGreeting  = "Hallo"
	ItalianGreeting = "Ciao"
	ChineseGreeting = "嗨。"
)

func main() {
	var language string
	fmt.Println("Выберите язык ", language)
	fmt.Scan(&language)

	switch language {
	case "en":
		fmt.Println(EnglishGreeting)
	case "sp":
		fmt.Println(SpanishGreeting)
	case "fr":
		fmt.Println(FrenchGreeting)
	case "ru":
		fmt.Println(RussianGreeting)
	case "gr":
		fmt.Println(GermanGreeting)
	case "it":
		fmt.Println(ItalianGreeting)
	case "ch":
		fmt.Println(ChineseGreeting)
	default:
		fmt.Println("Быдло, нет такого языка в мой базе. АХАхаххаха")

	}
}
