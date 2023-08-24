package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current number: ", num)
}

// Add adds together multiple numbers
func Add(nums ...int) int {
	value := 0
	for _, v := range nums {
		printNum(v)
		value += v
	}

	return value
}
