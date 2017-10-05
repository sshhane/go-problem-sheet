package main

import (
	"fmt"
	// "sort"
)

func main() {

	// vars
	count := 3
	var listOne[3]int
	// var listTwo[3]int

	// poulate first array
	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d:", i+1)
		fmt.Scanf("%d ", &listOne[i])
	}
}
