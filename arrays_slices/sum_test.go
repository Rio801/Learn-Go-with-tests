package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	wanted := 15

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 2, 4})
	wanted := []int{3, 7}

	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
