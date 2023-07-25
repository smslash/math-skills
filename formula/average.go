package formula

func Average(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	var sum float64 = 0
	for _, num := range numbers {
		sum += float64(num)
	}
	average := sum / float64(len(numbers))
	return int(average + 0.5)
}
