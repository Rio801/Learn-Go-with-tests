package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Testing for Hello, name return", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test for Hello World return", func(t *testing.T) {
		got := Hello("World")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q ,Wanted: %q", got, want)
	}
}
