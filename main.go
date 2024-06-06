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

// Reads the data.txt file and coverts them into slice of integers

//  Functions are called with the slice of integers as arguments

// Results of the called fuctions are printed
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, ` The number of arguments should be 2
		Usage : 
				go run . <program name> <name of file>
		`)
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "data.txt" {
			fileName, err := os.ReadFile(os.Args[1])
			if err != nil {

				fmt.Fprintln(os.Stderr, `Error Reading File`)
				return
			}
			if len(fileName) == 0 {
				fmt.Fprintln(os.Stderr, "The data file is empty")
				os.Exit(1)
			}
			var numberS []int
			split := strings.Split(string(fileName), "\n")
			for i := 0; i < len(split); i++ {
				if split[i] == "" {
					continue
				}
				numb, err := strconv.Atoi(split[i])
				if err != nil {
					Error, ok := err.(*strconv.NumError)
					if ok {
						if Error.Err == strconv.ErrSyntax {
							fmt.Fprintf(os.Stderr, `SyntaxError: Expecting integers, Check line: %v`, i+1)
							fmt.Println()
							os.Exit(1)
						} else if Error.Err == strconv.ErrRange {
							fmt.Fprintf(os.Stderr, `Integer out of range: ( Only integers <= 9223372036854775807 || integers>= -9223372036854775808 are allowed.)
					Check line %v in the data file
					`, i+1)
							os.Exit(1)
						}
					}
				}
				numberS = append(numberS, numb)
			}
			fmt.Printf("Average: %v\n", int(math.Round(average.Average(numberS))))
			fmt.Printf("Median: %v\n", int(math.Round(median.Median(numberS))))
			fmt.Printf("Variance: %v\n", int(math.Round(variance.Variance(numberS))))
			fmt.Printf("Standard Deviation: %v\n", int(math.Round(standard.Deviation(numberS))))
		} else {
			fmt.Fprintln(os.Stderr, `The expected second argument is data.txt
		Usage: 
			go run . [program name] <data.txt>`)
		}
	}
}
