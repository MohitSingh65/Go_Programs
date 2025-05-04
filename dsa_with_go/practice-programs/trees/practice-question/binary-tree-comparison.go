package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Head *Node
}

func compare(a *Node, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}
