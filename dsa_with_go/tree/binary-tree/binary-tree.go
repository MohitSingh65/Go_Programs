package main

import "fmt"

type TreeNode struct {
	Data  int       // value stored in the node
	Left  *TreeNode // Pointer to the left child
	Right *TreeNode // Pointer to the right child
}

// Insert inserts a new node with the given value into the binary tree
func Insert(root *TreeNode, data int) *TreeNode {
	if root == nil {
		// If the tree is empty, create a new node
		return &TreeNode{Data: data}
	}

	// Recursively find the correct position for the new node
	if data < root.Data {
		root.Left = Insert(root.Left, data)
	} else {
		root.Right = Insert(root.Right, data)
	}

	return root
}

func InOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)  // Traverse the left subtree
		fmt.Printf("%d", root.Data)  // Visit the root
		InOrderTraversal(root.Right) // Traverse the right subtree
	}
}

func PreOrderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d", root.Data)   // Visit the root
		PreOrderTraversal(root.Left)  // Traverse the left subtree
		PreOrderTraversal(root.Right) // Traverse the right subtree
	}
}

func PostOrderTraversal(root *TreeNode) {
	if root != nil {
		PostOrderTraversal(root.Left)  // Traverse the left subtree
		PostOrderTraversal(root.Right) // Traverse the right subtree
		fmt.Printf("%d", root.Data)    // Visit the root
	}
}

func Search(root *TreeNode, data int) bool {
	if root == nil {
		return false
	}

	if root.Data == data {
		return true
	}

	if data < root.Data {
		return Search(root.Left, data)
	}

	return Search(root.Right, data)
}

func main() {
	var root *TreeNode // Inititialize empty binary tree

	// Insert element into the binary tree
	root = Insert(root, 50)
	root = Insert(root, 30)
	root = Insert(root, 70)
	root = Insert(root, 20)
	root = Insert(root, 40)
	root = Insert(root, 60)
	root = Insert(root, 80)

	// Perform traversals
	fmt.Println("In-order Traversal:")
	InOrderTraversal(root)
	fmt.Println()

	fmt.Println("Pre-order Traversal:")
	PreOrderTraversal(root)
	fmt.Println()

	fmt.Println("Post-order Traversal:")
	PostOrderTraversal(root)
	fmt.Println()

	// Search for elements
	fmt.Println("Search for 40:", Search(root, 40))   // Should return true
	fmt.Println("Search for 100:", Search(root, 100)) // Should return false
}
