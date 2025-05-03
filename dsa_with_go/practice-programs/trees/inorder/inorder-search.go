package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	head *Node
}

func walk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	// pre

	// recurse
	walk(curr.Left, path)
	path = append(path, curr.Value)
	walk(curr.Right, path)

	// post
	return path
}

func in_order_traversal(head *Node) []int {

	return walk(head, []int{})
}
