package main

import "fmt"

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				tmp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = tmp
			}

		}
	}
}

func main() {
	numbers := []int{29, 11, 34, 45, 23, 101}
	bubbleSort(numbers)
	fmt.Print(numbers)
}
