package main

import (
	"fmt"
)

func linearSearch(numbers []int, searchNumber int) bool {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == searchNumber {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{12, 14, 27, 34, 89, 112}
	var searchNumber int

	fmt.Println("Enter a number to search: ")
	fmt.Scan(&searchNumber)

	found := linearSearch(numbers, searchNumber)
	if found {
		fmt.Printf("Number %v found!", searchNumber)
	} else {
		fmt.Printf("Number %v not found!", searchNumber)
	}

}
