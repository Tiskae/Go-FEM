package main

import (
	"GO-FEM/05_toolkit/code/utils"
	"fmt"
)

func calculateImportantData() int {
	value := utils.Add(1, 2, 3, 4, 5)
	fmt.Println(value)
	return value
}

func main() {
	result := calculateImportantData()
	fmt.Println(result)
}
