package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight float64
}

func hasUnvisited(seen []bool, dists []float64) bool {
	for i := 0; i < len(dists); i++ {
		if !seen[i] && !math.IsInf(dists[i], 1) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := 0
	lowestDistance := math.Inf(1)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}
	return idx
}

func dijkstra_list(source int, sink int, adjList [][]Edge) []int {
	arr := adjList
	n := len(arr)
	seen := make([]bool, n)
	dists := make([]float64, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dists[i] = math.Inf(1)
		prev[i] = -1
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}
	out := []int{}
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source)

	reverse(out)
	return out
}

func reverse(out []int) {
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
}

func main() {
	graph := [][]Edge{
		{{To: 1, Weight: 1}, {To: 2, Weight: 4}},
		{{To: 2, Weight: 2}, {To: 3, Weight: 6}},
		{{To: 3, Weight: 3}},
		{},
	}

	path := dijkstra_list(0, 3, graph)
	fmt.Println("Shortest path from 0 to 3:", path)
}
