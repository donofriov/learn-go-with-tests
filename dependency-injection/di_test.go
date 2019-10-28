package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vince")

	got := buffer.String()
	want := "Hello, Vince"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
