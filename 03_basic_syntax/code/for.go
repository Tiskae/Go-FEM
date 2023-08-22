// // Uncomment the entire file

package main

import "fmt"

func main() {

	// ****************************

	i := 1

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// ****************************

	j := 1

	for j <= 100 {
		fmt.Println(i)
		// This will behave like a while loop
		j += 1
	}

	// 	// ****************************

	var mySentence string = "This is a sentence"

	for index, value := range mySentence {
		fmt.Printf("Index: %d, Value: %s\n", index, string(value))
	}
}
