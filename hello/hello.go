package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns the first statement
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
