package median

func Median(numbers []int) float64 {
	var median float64

	currIndex := len(numbers) / 2 // Calculate the index for the middle element
	if len(numbers)%2 == 0 {      // Check if the number of data points is even
		median = (float64(numbers[currIndex]) + float64(numbers[currIndex-1])) / 2 // Calculate the median for even data points
	} else if len(numbers)%2 != 0 {
		median = (float64(numbers[currIndex])) // Calculate the median for odd data points
	}
	return median
}
