package median

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMedian(t *testing.T) {
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
	expected := 118
	found := int(math.Round(Median(numberS)))
	if found != expected {
		t.Errorf("Found: %v, want: %v", found, expected)
	}
}
