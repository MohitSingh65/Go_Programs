package main

import "fmt"

var maze = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

var dir = []Point{
	{0, 1},  // down
	{1, 0},  // right
	{0, -1}, // up
	{-1, 0}, // left
}

type Point struct {
	x, y int
}

// base case
// off the map
func walk(maze []string, wall byte, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// on a wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, curr)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for _, direction := range dir {
		next := Point{curr.x + direction.x, curr.y + direction.y}
		if walk(maze, wall, next, end, seen, path) {
			return true
		}
	}

	// Backtrack
	*path = (*path)[:len(*path)-1]
	return false

}

func main() {
	// Initialize seen matrix
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point
	start := Point{x: 10, y: 0} // First space after wall in top row
	end := Point{x: 1, y: 5}    // Last space before wall in bottom row

	if walk(maze, 'x', start, end, seen, &path) {
		fmt.Println("Path found:", path)
		// Visualize path in maze
		for _, p := range path {
			// Replace position with '*' in maze copy
			row := []rune(maze[p.y])
			row[p.x] = '*'
			maze[p.y] = string(row)
		}
		for _, row := range maze {
			fmt.Println(row)
		}
	} else {
		fmt.Println("No path found")
	}
}
