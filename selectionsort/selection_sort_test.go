package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSortReturnsArraySortedFromLowestToHighest(t *testing.T) {
	var unsorted = []int{5, 3, 6, 2, 10}

	var exp = []int{2, 3, 5, 6, 10}
	var got = SelectionSort(unsorted)

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %v to match %v", exp, got)
	}
}

func Test_findSmallestReturnsLocationOfSmallestInteger(t *testing.T) {
	var cases = []struct {
		exp int
		arr []int
	}{
		{3, []int{5, 3, 6, 2, 10}},
		{1, []int{6, 2, 10}},
		{2, []int{5, 3, 2, 6}},
		{0, []int{5}},
	}

	for _, c := range cases {
		var (
			exp = c.exp
			got = findSmallest(c.arr)
		)
		if exp != got {
			t.Errorf("expected %d, got %d; using %v", exp, got, c.arr)
		}
	}
}
