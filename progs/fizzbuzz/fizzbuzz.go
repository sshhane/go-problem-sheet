//fizzbuzz test
package main

import "fmt"

func main() {
	fmt.Println("Fizz Buzz")
	fmt.Println("=========\n")

	for i := 0; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("Fizz")
		}

		// else if i%5 == 0 {
		// 	fmt.Println("Buzz")
		// }

		// else if i%3 != 0 && i%5 != 0 {
		// 	fmt.Println(i)
		// 	return
		// }

		// else {
		// 	fmt.Println(i)
		// }

	}//for

}