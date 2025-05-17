package main

import "fmt"

func walk(adjMatrix [][]int, curr int, needle int, seen []bool, path *[]int) bool {
	graph := adjMatrix

	if seen[curr] {
		return false
	}

	seen[curr] = true

	*path = append(*path, curr)

	if curr == needle {
		return true
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := graph[curr][i]
		if edge == 0 {
			continue
		}
		if walk(graph, i, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func dfs(adjMatrix [][]int, source int, needle int) []int {
	graph := adjMatrix
	n := len(adjMatrix)
	seen := make([]bool, n)
	path := []int{}

	walk(graph, source, needle, seen, &path)

	if len(path) == 0 {
		return nil
	}
	return path
}

func main() {
	graph := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 1},
		{0, 1, 1, 0, 1},
		{0, 0, 1, 1, 0},
	}

	path := dfs(graph, 0, 4)
	fmt.Println("Path from 0 to 4:", path)
}
