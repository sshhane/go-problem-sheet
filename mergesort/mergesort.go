// http://austingwalters.com/merge-sort-in-go-golang/
// https://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang

package main

import (
	"fmt"
	// "sort"
)

// sort
// func sort(l, r []int) []int {


// }

// merge
func append(l, r []int) []int {

	// for r
	result := l.append(l, r...)
	return result
}

func main() {

	// vars
	count := 3
	var slice []int
	var listOne[3]int
	var listTwo[3]int

	// poulate arrays
	fmt.Println("First array")
	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d:", i+1)
		fmt.Scanf("%d ", &listOne[i])
	}

	fmt.Println("\nSecond array")
	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d:", i+1)
		fmt.Scanf("%d ", &listTwo[i])
	}


	for i := 0; i < count; i++ {
		// append slices
		slice = append(listOne[:], listTwo[:])
	}


	// call merge
	// fmt.Println("\n--- sorted ---\n\n", sort(listOne, listTwo), "\n")

}
