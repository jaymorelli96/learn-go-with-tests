package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jay", "")
		want := "Hello, Jay"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is passed as argument", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in English when an empty string is passed", func(t *testing.T) {
		got := Hello("Juan", "")
		want := "Hello, Juan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Joao", "Portuguese")
		want := "Ola, Joao"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
