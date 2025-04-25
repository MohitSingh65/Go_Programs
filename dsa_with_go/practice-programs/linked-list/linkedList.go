package main

type DoublyLinkedList struct {
	Length int
	head   *Node
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
	} else if idx == 0 {
		list.Prepend(item)
	}
	curr := list.head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	node := &Node{Value: item}
	node.Next = curr
	node.Prev = curr.Prev
	curr.Prev.Next = node
	curr.Prev = node
	list.Length++
}

func (list *DoublyLinkedList) Append(item int) {

}
