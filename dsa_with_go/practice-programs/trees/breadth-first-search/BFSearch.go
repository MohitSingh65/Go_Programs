package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Head *Node
}

func Bfs(head *Node, needle int) bool {
	q := []*Node{head}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			continue
		}

		if curr.Value == needle {
			return true
		}
		q = append(q, curr.Left)
		q = append(q, curr.Right)
	}

	return false
}

func main() {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	tree := Tree{Head: root}

	// Search for a value that exists
	found := Bfs(tree.Head, 4)
	fmt.Println("Found 4?", found) // true

	// Search for a value that does not exist
	found = Bfs(tree.Head, 99)
	fmt.Println("Found 99?", found) // false
}
