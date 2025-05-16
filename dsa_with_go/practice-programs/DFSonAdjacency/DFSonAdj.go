package main

import "fmt"

func bfs(adjMatrix [][]int, source int, needle int) []int {
	graph := adjMatrix
	n := len(adjMatrix)
	seen := make([]bool, n)
	prev := make([]int, n)

	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}
			seen[i] = true
			prev[i] = curr
			q = append(q, i)

		}
	}

	if prev[needle] == -1 {
		return nil
	}
	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, curr)

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}

func main() {
	adjMatrix := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 1},
		{0, 1, 1, 0, 1},
		{0, 0, 1, 1, 0},
	}

	source := 1
	needle := 4

	path := bfs(adjMatrix, source, needle)
	if path == nil {
		fmt.Printf("No path found from %d to %d\n", source, needle)
	} else {
		fmt.Printf("Path from %d to %d: %v\n", source, needle, path)
	}
}
