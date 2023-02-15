package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alo")

	got := buffer.String()
	want := "Hello, Alo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}