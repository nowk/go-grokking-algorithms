package binarysearch

func BinarySearch(list []int, item int) (int, int) {
	var (
		i     = -1
		steps = 0
	)

	l := 0
	h := len(list) - 1

	for l <= h {
		// count the # of steps taken to find the location, each iteration is 1
		steps++

		var (
			m     = l + h
			guess = list[m]
		)

		if guess == item {
			i = m
			break // if item is found
		}

		// calculate the h(igh)/l(ow) indexes
		if guess > item {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return i, steps
}
