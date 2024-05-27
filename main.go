package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/average"
	"math-skills/median"
	"math-skills/standard"
	"math-skills/variance"
)

func main() {
	// Check if the number of command-line arguments is 1
	if len(os.Args) == 1 {
		// Print a message indicating the correct number of arguments and usage instructions
		fmt.Fprintln(os.Stderr, ` The number of arguments should be 3
		Usage : 
				go run . <program name> <name of file: data.txt>
		`)
	}
	// Check if the number of command-line arguments is 2
	if len(os.Args) == 2 {
		// Check if the second argument is "data.txt"
		if os.Args[1] == "data.txt" {
			// Read the contents of the file named "data.txt"
			fileName, err := os.ReadFile(os.Args[1])
			if err != nil {
				// Print an error message if there is an issue reading the file
				fmt.Fprintln(os.Stderr, `Error Reading File`)
				return
			}
			// Check if the read file is empty
			if len(fileName) == 0 {
				fmt.Fprintln(os.Stderr, "The data file is empty")
				os.Exit(1)
			}
			var numberS []int
			// Split the file content by newline character
			split := strings.Split(string(fileName), "\n")
			for i := 0; i < len(split); i++ {
				if split[i] == "" {
					continue
				}
				// Convert the string to an integer
				numb, err := strconv.Atoi(split[i])
				// Check for errors during conversion and handle them accordingly
				if err != nil {
					Error, ok := err.(*strconv.NumError)
					if ok {
						if Error.Err == strconv.ErrSyntax {
							// Print a syntax error message with the line number
							fmt.Fprintf(os.Stderr, `SyntaxError: Expecting integers, Check line: %v`, i+1)
							fmt.Println()
							os.Exit(1)
						} else if Error.Err == strconv.ErrRange {
							// Print an integer out of range error message with the line number
							fmt.Fprintf(os.Stderr, `Integer out of range: ( Only integers <= 9223372036854775807 || integers>= -9223372036854775808 are allowed.)
					Check line %v in the data file
					`, i+1)
							os.Exit(1)
						}
					}
				}
				// Append the integer to the numberS slice
				numberS = append(numberS, numb)
			}
			// Calculate and print the average, median, variance, and standard deviation of the numbers
			fmt.Printf("Average: %v\n", int(math.Round(average.Average(numberS))))
			fmt.Printf("Median: %v\n", int(math.Round(median.Median(numberS))))
			fmt.Printf("Variance: %v\n", int(math.Round(variance.Variance(numberS))))
			fmt.Printf("Standard Deviation: %v\n", int(math.Round(standard.Deviation(numberS))))
		} else {
			// Print a message indicating the expected third argument and usage instructions
			fmt.Fprintln(os.Stderr, `The expected third argument should be data.txt
		Usage: 
			go run . [program name] <data file name>`)
		}
	}
}
