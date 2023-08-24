package main

import "fmt"

func getAverage(scores []float64) (average float64) {
	sum := 0.0
	for _, value := range scores {
		sum += value
	}
	average = sum / float64(len(scores))
	return
}

func isKeyInMap(the_map map[string]string, key string) (hasKey bool) {
	_, hasKey = the_map[key]
	return
}

func appendToSlice(the_slice []string, items ...string) (updatedSlice []string) {
	// Spread operator 🚀
	updatedSlice = append(the_slice, items...)
	fmt.Println(updatedSlice)
	return
}

func main() {
	// Part 1
	scores := []float64{77, 82, 83, 74, 60}
	average := getAverage(scores)
	fmt.Println(scores, average)

	// Part 2
	myCars := map[string]string{
		"daily":    "BMW M4",
		"drag":     "Dodge Challenger Demon SRT",
		"luxury":   "Bentley Flying Spur",
		"supercar": "Lamborghini Aventador SVJ",
		"hypercar": "Pagani Zonda R",
	}

	if has_key := isKeyInMap(myCars, "offroad"); has_key {
		fmt.Println("Key present")
	} else {
		fmt.Println("Key absent")
	}

	// Part 3
	groceries := []string{"Apples", "Conflakes", "Liquid soap", "Toothpaste"}
	appendToSlice(groceries, "Cassava", "Garri", "Ewa", "Agbado")
}
