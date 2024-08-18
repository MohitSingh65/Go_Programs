package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}

func main() {
	num1 := 5
	num2 := 3

	sum := addition(num1, num2)
	fmt.Println("The sum is:", sum)
}