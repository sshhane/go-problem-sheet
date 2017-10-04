// largest & smallest in list
package main
import "fmt"

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

	// print array
}