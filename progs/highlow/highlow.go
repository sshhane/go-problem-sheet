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
} //genRand

func main() {
	fmt.Println("High Low, 世界\n")

	// vars
    randNum := genRand(10)	// populate var w/ rand num
    guesses := 0
    var guess int
    guessLimit := 5


    // for guess limit
    for guesses < guessLimit {
        //ask user for guess
        fmt.Println("Guess a number between 1 and 10: ")
        fmt.Scanf("%d ", &guess)
        guesses++
        // if guess is too high or low
        if guess < randNum {
            fmt.Println("Too low!")
        } else if guess > randNum {
            fmt.Println("Too high!")
        } else if guess == randNum {
            fmt.Printf("Correct!, number was: %d\n", randNum)
            break
        } else {
        	fmt.Println("Out of Guesses ):")
            break
        }
	}

}

    // fmt.Println("Random num:")
    // fmt.Println(random.Intn(100))	// generate random num