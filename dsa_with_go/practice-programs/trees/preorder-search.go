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
	path = append(path, curr.Value)

	// recurse
	walk(curr.Left, path)
	walk(curr.Right, path)

	// post
	return path
}

func Pre_order_traversal(head *Node) []int {

}
