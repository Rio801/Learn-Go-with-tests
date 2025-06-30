package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictonary{"test": "this is just a test"}

	got := dictionary.SearchDic("test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got: %q ,want:%q", got, want)
	}

}
