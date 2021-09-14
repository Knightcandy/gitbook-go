package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Greeting("World", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World when empty string supplied", func(t *testing.T) {
		got := Greeting("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Greeting("Elode", "Spanish")
		want := "Hola Elode"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Greeting("Elode", "French")
		want := "Bonjour Elode"

		assertCorrectMessage(t, got, want)
	})
}
