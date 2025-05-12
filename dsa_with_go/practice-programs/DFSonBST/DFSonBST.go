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

func search(curr *Node, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		return search(curr.Right, needle)
	}
	return search(curr.Left, needle)
}

func main() {
	// Create nodes manually for a BST
	root := &Node{Value: 10}
	root.Left = &Node{Value: 5}
	root.Right = &Node{Value: 15}
	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 7}
	root.Right.Left = &Node{Value: 12}
	root.Right.Right = &Node{Value: 17}

	// Create tree
	tree := Tree{Head: root}

	// Test values
	testVals := []int{7, 12, 20, 1}
	for _, val := range testVals {
		found := search(tree.Head, val)
		if found {
			fmt.Printf("Value %d found in tree.\n", val)
		} else {
			fmt.Printf("Value %d NOT found in tree.\n", val)
		}
	}
}
