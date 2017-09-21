//fizzbuzz test
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fizz Buzz")
	fmt.Println("=========\n")

	for i := 0; i <= 100; i++ {


		if( i%3 == 0 && i%5 == 0 ) {
		fmt.Println("FizzBuzz");
		} else if i%5 == 0 {
		fmt.Println("Buzz");
		} else {
		fmt.Printf("%d\n", i);
		}

		// if (i%3 == 0 && i%5 == 0) {
		// 	fmt.Println("fizzbuzz")
		// }

		// else if i%3 == 0 {
		// 	fmt.Println("Fizz")
		// }

		// else if i%5 == 0 {
		// 	fmt.Println("Buzz")
		// }

		// else {
		// 	fmt.Println(i)
		// }

	}//for

}