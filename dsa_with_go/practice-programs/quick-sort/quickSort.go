package main

import "fmt"

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[idx], arr[hi] = arr[hi], arr[idx]

	return idx
}

func main() {
	arr := []int{8, 3, 5, 12, 9, 6, 10}
	fmt.Println("Before: ", arr)
	qs(arr, 0, len(arr)-1)
	fmt.Println("After: ", arr)

}
