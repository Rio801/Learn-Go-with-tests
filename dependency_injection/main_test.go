package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Test Greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()

		want := "Hello, Chris \n"

		if got != want {
			t.Errorf("got: %q ,want: %q", got, want)
		}

	})

	t.Run("Countdown", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := &SpySleeper{}
		Countdown(&buffer, sleeper)
		got := buffer.String()
		want := "3 \n2 \n1 \nGo!"

		if got != want {
			t.Errorf("got: %q ,want: %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("wnat 3 calls got:%d", sleeper)
		}
	})
}
