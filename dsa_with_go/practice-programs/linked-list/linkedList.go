package main

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
	curr.Prev.Next = node
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
