package main

import "fmt"

func main() {
	// All are equivalent
	// var scores = [5]int{265, 378, 554, 654, 238}
	// var scores [5]int = [5]int{265, 378, 554, 654, 238}
	// scores := [5]int{265, 378, 554, 654, 238}
	scores := [...]int{265, 378, 554, 654, 238}
	// scores[0] = 265
	// scores[1] = 378
	// scores[2] = 554
	// scores[3] = 654
	// scores[4] = 238

	for _, value := range scores {
		fmt.Println(value)
	}

	fmt.Println(scores)
}
