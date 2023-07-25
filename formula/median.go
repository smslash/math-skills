package formula

import "sort"

func Median(numbers []int) int {
	length := len(numbers)

	if length == 0 {
		return 0
	}

	sorted := make([]int, length)
	copy(sorted, numbers)
	sort.Ints(sorted)

	var median float64
	if length%2 == 0 {
		median = (float64(sorted[length/2-1]) + float64(sorted[length/2])) / 2
	} else {
		median = float64(sorted[length/2])
	}

	return int(median + 0.5)
}
