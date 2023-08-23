package main

import "fmt"

func getSum(nums []float64) (sum float64) {
	sum = 0

	for _, value := range nums {
		sum += value
	}

	return
}

func getAverage(nums ...float64) (average float64) {
	sum := getSum(nums)
	length := float64(len(nums))
	average = sum / length
	return
}

func main() {
	average := getAverage(3.99, 4.33, 4.56, 4.78, 5.00, 4.44, 4.53)
	fmt.Printf("The average of all the numbers is: %f", average)
}
