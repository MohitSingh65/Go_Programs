package main

import (
	"fmt"
)

func main() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4

	for i, v := range numbers {
		fmt.Printf("%s:%d", i, v)
	}
}
