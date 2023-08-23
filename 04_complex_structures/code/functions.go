package main

import "fmt"

func printAge(age1, age2 int) (ageOfSally, ageOfBob int) {
	ageOfSally = age1
	ageOfBob = age2
	return
}

func printUnlimitedAges(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
}

func dummyFunc(num int) (int, int) {
	num1 := num
	return num1, 0
}

func main() {
	x, y := printAge(10, 21)
	fmt.Println(x)
	fmt.Println(y)
	printUnlimitedAges(21, 18, 23, 49)
}
