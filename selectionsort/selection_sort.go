package selectionsort

func SelectionSort(in []int) []int {
	var out = []int{}

	for len(in) > 0 {
		var i = findSmallest(in)
		out = append(out, in[i])

		// remove smallest
		in = append(in[:i], in[i+1:]...)
	}

	return out
}

// findSmallest fines the location of the smallest int in an array of ints
func findSmallest(in []int) int {
	var smallest = in[0]
	var i = 0

	for n, v := range in[1:] {
		if v < smallest {
			smallest = v
			i = n + 1
		}
	}

	return i
}
