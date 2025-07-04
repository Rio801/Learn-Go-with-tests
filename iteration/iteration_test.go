package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q, Got: %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a")
	}
}
