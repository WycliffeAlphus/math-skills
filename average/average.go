package average

func Average(numbers []int) float64 {
	var mean float64  // Initialize the mean variable to store the average
	var total float64 // Initialize the total variable to store the sum of numbers

	for _, n := range numbers { // Iterate over each number in the slice
		total += float64(n)                  // Add the current number to the total sum
		mean = total / float64(len(numbers)) // Calculate the mean by dividing the total sum by the number of elements
	}
	return mean // Return the calculated mean
}
