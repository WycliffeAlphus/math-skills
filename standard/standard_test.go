package standard

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDeviation(t *testing.T) {
	var numberS []int
	fileName, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error Reading File`)
		os.Exit(1)
	}
	split := strings.Split(string(fileName), "\n")
	for i := 0; i < len(split); i++ {
		numb, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println("Error splitting number")
		}
		numberS = append(numberS, numb)
	}

	expected := 28

	found := int(float64(Deviation(numberS)))
	if found != expected {
		t.Errorf("Found: %v, want: %v", found, expected)
	}
}
