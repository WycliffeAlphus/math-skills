package median

func Median(numbers []int) float64 {
	var median float64

	for j := 0; j < len(numbers); j++ { // Loop through the numbers slice
		for i := 0; i < len(numbers); i++ { // Nested loop to compare adjacent elements
			if (i+1) < len(numbers) && numbers[i+1] < numbers[i] { // Check if the next element is smaller
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i] // Swap the elements if needed
			}
		}
	}

	currIndex := len(numbers) / 2 // Calculate the index for the middle element
	if len(numbers)%2 == 0 {      // Check if the number of data points is even
		median = (float64(numbers[currIndex]) + float64(numbers[currIndex-1])) / 2 // Calculate the median for even data points
	} else if len(numbers)%2 != 0 {
		median = (float64(numbers[currIndex])) // Calculate the median for odd data points
	}
	return median
}
