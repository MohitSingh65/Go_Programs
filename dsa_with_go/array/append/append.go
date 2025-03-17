package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums = append([]int{0}, nums...) // Add new element to the start
	nums = append(nums, 4)           // Add new element to the end
	fmt.Println(nums)
	nums = nums[:len(nums)-1] // Removes the last element
	fmt.Println(nums)
}
