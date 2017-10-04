// HighLow Game
package main

import "time"
import "fmt"
import "math/rand"

func main() {
	fmt.Println("High Low, 世界\n")

	//Generate number
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Println("Random num:")
    fmt.Println(r1.Intn(100))
    fmt.Println(r1.Intn(100))
    
}