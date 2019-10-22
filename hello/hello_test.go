package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Vince")
	want := "Hello, Vince!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
