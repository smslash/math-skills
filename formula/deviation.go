package formula

import (
	"math"
)

func StandardDeviation(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var sum float64
	for _, v := range nums {
		sum += float64(v)
	}
	mean := sum / float64(len(nums))

	var variance float64
	for _, v := range nums {
		diff := float64(v) - mean
		variance += diff * diff
	}
	variance /= float64(len(nums))
	stdDev := math.Sqrt(variance)

	return int(math.Round(stdDev))
}
