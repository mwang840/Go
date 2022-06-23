package src

//Will be implementing a swap function (Bubble Sort)
import("fmt"
 		"sort"
)
//Main function to test swapped values
func sortNumbers(numbers []int) []int {
	var sorted[]int
	if(numbers == NULL){
		return nil
	}else {
		sorted = sort.Ints(numbers)
		return sorted
	}
}

func main(){
	fmtPrintln(sortNumbers([1, 2, 3, 10, 5]))
}
