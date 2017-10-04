// HighLow Game
package main

import(
    "fmt"
    "math/rand"
    "time"
)

// Generate number
func genRand(max int) int{
	source := rand.NewSource(time.Now().UnixNano())	//seed for number is current time
    random := rand.New(source)
    randNum := random.Intn(max)
	return randNum
}

func main() {
	fmt.Println("High Low, 世界\n")


	// vars
	
    // populate var w/ rand num
    randNum := genRand(10)

    // start game

    fmt.Println("Guess a number between 1 & 10")

    fmt.Println(randNum)

    // fmt.Println("Random num:")
    // fmt.Println(random.Intn(100))	// generate random num
    
}