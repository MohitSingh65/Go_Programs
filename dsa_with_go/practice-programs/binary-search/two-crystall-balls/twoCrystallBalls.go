package main

import (
	"fmt"
	"math"
)

func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))
	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount
	for ; i < len(breaks); i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}

func main() {
	breaks := []bool{false, false, false, false, true, true, true}
	fmt.Println(twoCrystalBalls(breaks))
}
