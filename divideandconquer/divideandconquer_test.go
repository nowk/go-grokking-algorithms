package divideandconquer

import (
	"testing"
)

func TestSum(t *testing.T) {
	var exp = 12
	var got = Sum([]int{2, 4, 6})

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestSumr(t *testing.T) {
	var exp = 12
	var got = Sumr([]int{2, 4, 6})

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
