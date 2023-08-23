package main

import "fmt"

func getSum(n1, n2, n3 float64) float64 {
	return n1 + n2 + n3
}

func getAverage(num1, num2, num3 float64) (average float64) {
	sum := getSum(num1, num2, num3)
	average = sum / 3
	return
}

func main() {
	average := getAverage(4.1, 3.99, 5.00)
	fmt.Printf("The average of the numbers is: %f", average)
}
