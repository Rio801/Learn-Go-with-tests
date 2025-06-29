package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	result := 4

	t.Run("Adding two intergers", func(t *testing.T) {
		assertCorrectMessage(t, sum, result)
	})
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %d , Expected : %d", got, want)
	}
}
