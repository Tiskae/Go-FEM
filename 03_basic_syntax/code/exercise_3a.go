package main

import "fmt"

func main() {
	mySentence := "Is this really a sentence? Oh! it is."

	for index, value := range mySentence {
		if index%2 == 1 {
			fmt.Printf("Index: %d, value: %s\n", index, string(value))
		}
	}
}
