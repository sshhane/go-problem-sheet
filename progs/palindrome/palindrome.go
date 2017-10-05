// palindrome test
// http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (
	"fmt"
	"strings"
)

// func isPalindrome
func isPalindrome(str string) string {

	m := len(str) / 2	// middle char
	l := len(str) - 1	// last char

	for i:=0; i<m; i++ {	// iterate until middle char
		if str[i] != str[l-i] {	// is 'first' char != 'last'
			return "Not Palindrome"
		}
	} // for
	// is pal
	return "Is Palindrome"

} // isPalindrome

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
	fmt.Println(isPalindrome(myStr))
}