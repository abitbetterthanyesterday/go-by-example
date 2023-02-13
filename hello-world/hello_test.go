package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Alo", "")
		want := "Hello, Alo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Alo", "Spanish")
		want := "Hola, Alo"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Alo", "French")
		want := "Bonjour, Alo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
