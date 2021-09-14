package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const greet = "Hello "
const spanishGreet = "Hola "
const frenchGreet = "Bonjour "

func Greeting(name string, lang string) string {
	if name == "" {
		return greet + "World"
	}

	return getGreeting(lang) + name
}

func getGreeting(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishGreet
	case french:
		prefix = frenchGreet
	default:
		prefix = greet
	}
	return
}

func main() {
	fmt.Println(Greeting("World", ""))
}
