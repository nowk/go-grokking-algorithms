package divideandconquer

// Sum a basic loop to sum an array
func Sum(a []int) int {
	var total = 0

	for _, v := range a {
		total += v
	}

	return total
}

// Sumr is a recursive version of Sum
func Sumr(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + Sumr(a[1:])
}
