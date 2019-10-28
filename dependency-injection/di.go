package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet function
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler function
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world!")
}

// Main function
func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
