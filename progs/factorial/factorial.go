//factorial
package main
import "fmt"
 
/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
                       // Range: 0 through 18446744073709551615. 
var i int = 1
var n int
 
/*     function declaration        */
func factorial(n int) uint64 {   

	// check if negative
    if(n < 0){
        fmt.Print("Please use a positive number:")   

    }else{        
        for i:=1; i<=n; i++ {
            factVal *= uint64(i)  // mismatched types int64 and int
        }
         
    }    
    return factVal  /* return from function*/
}//factorial

// func sum(var factVal int){
// 	var sum int = 0
// 	sum = factVal + 1
// 	return sum
// }//sum

func sumOfDigits(factVal uint64) uint64{

    // String digits = new Integer(factVal).toString();
    var sum uint64 = 0
    sum = factVal + 1
    // int sum = 0;
    // for (char c: digits.toCharArray())
    //     sum += c - '0';
    return sum;
}
 
func main(){    
    fmt.Print("Enter a positive integer between 0 - 50 : ")
    fmt.Scan(&n)
    fmt.Print("Factorial is: ",factorial(n))
	// fmt.Print("Factorial is: ",factVal)
	fmt.Print("Sum is: ",sumOfDigits(factVal))

}//main