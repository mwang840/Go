package src

//Will be implementing a swap function (Bubble Sort)
import (
	"sort"
)

// Main function to test swapped values
func sortNumbers(numbers []int) []int {
	var sorted []int
	if numbers == nil {
		return nil
	} else {
		sorted = sort.Ints(numbers)
		return sorted
	}
}

func main() {
	fmtPrintln()
}
