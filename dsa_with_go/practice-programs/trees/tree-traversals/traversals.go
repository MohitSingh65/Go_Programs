package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func preOrderWalk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	path = append(path, curr.Value)
	path = preOrderWalk(curr.Left, path)
	path = preOrderWalk(curr.Right, path)
	return path
}

func inOrderWalk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	path = inOrderWalk(curr.Left, path)
	path = append(path, curr.Value)
	path = inOrderWalk(curr.Right, path)
	return path
}

func postOrderWalk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	path = postOrderWalk(curr.Left, path)
	path = postOrderWalk(curr.Right, path)
	path = append(path, curr.Value)
	return path
}

func PreOrderTraversal(root *Node) []int {
	return preOrderWalk(root, []int{})
}

func InOrderTraversal(root *Node) []int {
	return inOrderWalk(root, []int{})
}

func PostOrderTraversal(root *Node) []int {
	return postOrderWalk(root, []int{})
}

func main() {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Println("Pre-order traversal:  ", PreOrderTraversal(root))  // [1 2 4 5 3]
	fmt.Println("In-order traversal:   ", InOrderTraversal(root))   // [4 2 5 1 3]
	fmt.Println("Post-order traversal: ", PostOrderTraversal(root)) // [4 5 2 3 1]
}
