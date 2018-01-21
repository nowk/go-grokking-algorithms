package binarysearch

import (
	"testing"
)

func TestBinarySearchReturnsBothLocationAndStepsTaken(t *testing.T) {
	var cases = []struct {
		list  []int
		item  int
		i     int
		steps int
	}{
		{
			[]int{1, 4, 8, 12},
			8,
			2,
			2,
		},
		{
			[]int{1, 4, 8, 9, 12},
			8,
			2,
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			7,
			6,
			4,
		},

		// reason for 5 steps, is this triggers the `h = m - 1`, at start m
		// will be the "end of the list" and will work it's way backwards
		{
			[]int{1, 2, 3, 4, 5},
			0,
			-1,
			5,
		},

		// reason for 1 steps, is this triggers the `l = m + 1`, at start
		// making l > h thus ending the loop (search for more)
		{
			[]int{1, 2, 3, 4, 5},
			6,
			-1,
			1,
		},
	}

	for n, c := range cases {
		i, steps := BinarySearch(c.list, c.item)

		{
			var exp = c.i
			var got = i

			if exp != got {
				t.Errorf("case #%d; expected location to be %d, got %d", n, exp, got)
			}
		}

		{
			var exp = c.steps
			var got = steps

			if exp != got {
				t.Errorf("case #%d; expected %d steps taken, got %d", n, exp, got)
			}
		}
	}
}
