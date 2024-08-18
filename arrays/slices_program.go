package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	// append function
	s1 = append(s1, 4, 5)
	s2 := make([]int, 2)
	// copy function
	copy(s1, s2)
	fmt.Println(s1, s2)
}