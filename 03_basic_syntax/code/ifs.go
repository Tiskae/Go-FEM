// // Uncomment this entire file

package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func someFunction() error {
	return errors.New("some error")
}

func main() {

	var totalScore = 9

	if totalScore > 10 {
		fmt.Println(totalScore)
	}

	// ****************************

	if totalScore > 100 {
		fmt.Println("Greater than 100")
	} else if totalScore == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

	// ****************************
	err := someFunction()
	// => If this function returns a value,
	// => it will be an  error of type Error

	// ****************************
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := someFunction(); err != nil {
		fmt.Println(err.Error())
	}

	// End of file curly brace
	getScore()
}

func getScore() {
	var totalScore = int(math.Ceil(rand.Float64() * 100))

	fmt.Printf("Your score: %d\n", totalScore)

	if totalScore > 69 {
		fmt.Println("That's an A!")
	}

	if totalScore > 69 {
		fmt.Println("I repeat, that's an A!")
	} else if totalScore > 59 {
		fmt.Println("That's a B")
	} else if totalScore > 49 {
		fmt.Println("Hmm, that's a C")
	} else if totalScore > 44 {
		fmt.Println("Oh, that's a D")
	} else if totalScore > 39 {
		fmt.Println("Ah, that's an E")
	} else {
		fmt.Println("Omo!")
	}
}
