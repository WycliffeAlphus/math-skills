package median

// Returns middle numbers in data set

// Data sets have to be sorted

// For even data sets, find the average of the two middle numbers
func Median(numbers []int) float64 {
	var median float64

	if len(numbers) < 1 {
		return 0
	}

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers); i++ {
			if (i+1) < len(numbers) && numbers[i+1] < numbers[i] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}

	currIndex := len(numbers) / 2
	if len(numbers)%2 == 0 {
		median = (float64(numbers[currIndex]) + float64(numbers[currIndex-1])) / 2
	} else if len(numbers)%2 != 0 {
		median = (float64(numbers[currIndex]))
	}
	return median
}
