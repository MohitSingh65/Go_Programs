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
	walk(curr.Right, path)

	// post
	path = append(path, curr.Value)
	return path
}

func post_order_traversal(head *Node) []int {

	return walk(head, []int{})
}
