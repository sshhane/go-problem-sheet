// reverse a string
// https://coderwall.com/p/fw6miw/reverse-text-in-golang

package main
import "fmt"

func Reverse(s string) string {
	var result string
	//end of s to start
    for i := len(s)-1; i >= 0; i-- {
        result += string(s[i])	// populate result with chars in element
    }
    return result 
}

//main function 
func main() {
	//vars
	var str string
	//prompt
	fmt.Println("Enter string: ")
	fmt.Scanf("%s ", &str)
	
	//output result
	fmt.Printf("\nReverse of %s: %s", str, Reverse(str))
}