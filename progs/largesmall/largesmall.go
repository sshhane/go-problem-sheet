// largest & smallest in list
package main
import "fmt"

// func minMax(numbers [10]int) int, int {

// 	// init min / max
// 	var min int = numbers[0]
// 	var max int = numbers[0]

// 	// iterate overf array


// 	return min, max
// }

//function to produce min and max values
func minMax(array [10]int) (int, int) {

	//initialise min / max
	var min int = array[0]
	var max int = array[0]

	//iterate over array
    // for...
    return min, max
}

func main() {
	fmt.Println("Largest / Smallest")

	// vars
	size := 10
	var numbers [10]int

	fmt.Println("Enter 10 numbers:")

	// pop array
	for i:=0; i<size; i++ {

		// ask for values
		fmt.Printf("Enter number %d:", i+1)
		fmt.Scanf("%d ", &numbers[i])

	}

	// print min max
	fmt.Printf("Min and max are: ")
	fmt.Println(minMax(numbers))
}