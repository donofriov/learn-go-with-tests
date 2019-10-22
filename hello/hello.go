package main

import "fmt"

const french = "French"
const german = "German"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Servus, "
const spanishHelloPrefix = "Hola, "

// Hello returns the first statement
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
