package variance

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestVariance(t *testing.T) {
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
	expected := 785
	found := int(math.Round(Variance(numberS)))
	if found != expected {
		t.Errorf("Expected: %v, found %v ", expected, found)
	}
}
