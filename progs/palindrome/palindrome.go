// palindrome test
// http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (
	"fmt"
	"strings"
)

// func isPalindrome

func main() {
	fmt.Println("Palindrome test")

	// vars
	var myStr string

	// ask for string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &myStr)

	// toLowercase
	myStr = strings.ToLower(myStr)

	// call isPalindrome
	// fmt.Println(isPalindrome(myStr))
}