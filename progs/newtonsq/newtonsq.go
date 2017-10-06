package main
//import packages
import (
	"fmt"
	"math"
)

// constants
const DELTA = .0000001

func sRoot(x float64, z float64) float64{

    // result := func() float64 {
    //     return z - (z*z - x) / (2 * z)
}
}

func main() {
    //vars
    var num
    var start float64

    //prompt input
    fmt.Print("Enter num: ")
    fmt.Scanf("%f ", &num)

    fmt.Print("Pick starting point: ")
    fmt.Scanf("%f ", &start)

    fmt.Println("\nSquare Root: ", sqRoot(num, start))
}