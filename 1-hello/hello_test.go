package main

import "testing"

func TestHello(t *testing.T) {
	// assertCorrectMessage helper function
	assertCorrectMessage := func(t *testing.T, got, want string) {
		/*
		 t.Helper() tells the test suite that this function is a helper function
		*/
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Test case that test when Hello does not have an empty string
	t.Run("saying go-hello-with-test to people", func(t *testing.T) {
		got := Hello("Pascal", "")
		want := "Hello, Pascal"
		assertCorrectMessage(t, got, want)
	})

	// Test case that test when Hello has an empty string
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying go-hello-with-test in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying go-hello-with-test in French", func(t *testing.T) {
		got := Hello("Julien", "French")
		want := "Bonjour, Julien"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying go-hello-with-test in Haitian", func(t *testing.T) {
		got := Hello("Simeon", "Haitian")
		want := "Bonjou, Simeon"
		assertCorrectMessage(t, got, want)
	})
}
