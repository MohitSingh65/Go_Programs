package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func Prepend(head **Node, value int) {
	newNode := &Node{Value: value, Next: *head}
	*head = newNode
}

func printForward(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Value)
		current = current.Next
	}
}

func printBackward(tail *Node) {
	current := tail
	for current != nil {
		fmt.Printf("%d", current.Value)
		current = current.Prev
	}
}

func insertAfter(node *Node, value int) {
	newNode := &Node{Value: value}

	newNode.Next = node.Next
	newNode.Prev = node

	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode
}

func deleteNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

func main() {
	head := &Node{Value: 1}
	second := &Node{Value: 2}
	third := &Node{Value: 3}

	head.Next = second
	second.Next = third

	third.Prev = second
	second.Prev = head

	fmt.Println("Forward traversal:")
	printForward(head)

	fmt.Println("\nBackward traversal:")
	printBackward(third)

	insertAfter(second, 4)
	fmt.Println("\nAfter insertion:")
	printForward(head)

	deleteNode(second)
	fmt.Println("\nAfter deletion:")
	printForward(head)
}
