package main

import (
	"fmt"
)

type DoublyLinkedList struct {
	Length int
	head   *Node
	tail   *Node
}

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) Prepend(item int) {
	node := &Node{Value: item}
	list.Length++
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	node.Next = list.head
	list.head.Prev = node
	list.head = node
}

func (list *DoublyLinkedList) InsertAt(item int, idx int) {
	if idx > list.Length {
		panic("Oh no!")
	} else if idx == list.Length {
		list.Append(item)
		return
	} else if idx == 0 {
		list.Prepend(item)
		return
	}
	curr := list.head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	node := &Node{Value: item}
	node.Next = curr
	node.Prev = curr.Prev
	node.Prev.Next = node
	curr.Prev = node
	list.Length++
}

func (list *DoublyLinkedList) Append(item int) {
	list.Length++
	node := &Node{Value: item}

	if list.tail == nil {
		list.head = node
		list.tail = node
		return
	}
	node.Prev = list.tail
	list.tail.Next = node
	list.tail = node
}

func (list *DoublyLinkedList) Remove(item int) int {
	curr := list.head
	for i := 0; i < list.Length; i++ {
		if curr.Value == item {
			break
		}
		curr = curr.Next
	}
	if curr == nil {
		return 0
	}
	list.Length--

	if list.Length == 0 {
		out := list.head.Value
		list.head = nil
		list.tail = nil
		return out
	}

	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	}

	if curr == list.head {
		list.head = curr.Next
	}

	if curr == list.tail {
		list.tail = curr.Prev
	}

	removedValue := curr.Value
	return removedValue
}

func (list *DoublyLinkedList) Get(idx int) int {
	node := list.getAt(idx)
	if node == nil {
		panic("Index out of bounds")
	}
	return node.Value
}

func (list *DoublyLinkedList) RemoveAt(idx int) {
	node := list.getAt(idx)
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		list.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		list.tail = node.Prev
	}

	list.Length--
}

func (list *DoublyLinkedList) getAt(idx int) *Node {
	if idx < 0 || idx >= list.Length {
		return nil
	}
	curr := list.head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	return curr
}

func (list *DoublyLinkedList) Print() {
	fmt.Print("List: ")
	for curr := list.head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Value)
	}
	fmt.Println()
}

func (list *DoublyLinkedList) PrintReverse() {
	fmt.Print("Reverse: ")
	for curr := list.tail; curr != nil; curr = curr.Prev {
		fmt.Printf("%d ", curr.Value)
	}
	fmt.Println()
}

func main() {
	list := NewDoublyLinkedList()

	// Append values
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("After appending 10, 20, 30:")
	list.Print()        // List: 10 20 30
	list.PrintReverse() // Reverse: 30 20 10

	// Prepend a value
	list.Prepend(5)
	fmt.Println("\nAfter prepending 5:")
	list.Print()        // List: 5 10 20 30
	list.PrintReverse() // Reverse: 30 20 10 5

	// Insert at position
	list.InsertAt(15, 2)
	fmt.Println("\nAfter inserting 15 at index 2:")
	list.Print()        // List: 5 10 15 20 30
	list.PrintReverse() // Reverse: 30 20 15 10 5

	// Remove a value
	removed := list.Remove(20)
	fmt.Printf("\nRemoved: %d\n", removed)
	list.Print() // List: 5 10 15 30

	// Get value at index
	value := list.Get(2)
	fmt.Printf("\nValue at index 2: %d\n", value) // 15

	// Remove at index
	list.RemoveAt(0)
	fmt.Println("\nAfter removing index 0:")
	list.Print()        // List: 10 15 30
	list.PrintReverse() // Reverse: 30 15 10

	fmt.Printf("\nFinal Length: %d\n", list.Length) // Final Length: 3
}
