package main

import "fmt"

// Hello returns the first statement
func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
