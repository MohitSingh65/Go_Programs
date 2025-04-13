package main

import (
	"fmt"
)

func binarySearch(numbers []int, searchNumber int) bool {
	low, high := 0, len(numbers)-1

	for low < high {
		mid := low + (high-low)/2
		midValue := numbers[mid]

		if midValue == searchNumber {
			return true
		} else if midValue > searchNumber {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	numbers := []int{12, 14, 27, 34, 89, 112}
	var searchNumber int

	fmt.Println("Enter a number to search: ")
	fmt.Scan(&searchNumber)

	found := binarySearch(numbers, searchNumber)
	if found {
		fmt.Printf("Number %v found!", searchNumber)
	} else {
		fmt.Printf("Number %v not found.", searchNumber)
	}
}
