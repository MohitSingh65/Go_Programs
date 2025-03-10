package main

// Define Node structure
type node struct {
	val  int
	next *node
}

// Define the linked list structure
type LinkedList struct {
	head *node
}

// Insert node at beginning
func (ll *LinkedList) InsertAtStart(val int) {
	newNode := &Node{val: val, next: ll.head}
	ll.head = newNode
}

// Insert node at end
func (ll *LinkedList) InsertAtEnd(val int) {
	newNode := &Node{val: val, next: nil}

	if ll.head == nil {
		current = current.next
	}
	current.next
}

// Delete node with specific value
func (ll *LinkedList) Delete(val int) {
	if ll.head == nil {
		return
	}

	// Delete node at start
	if ll.head.val == val {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.val == val {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// Print the linked list
func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {

	}
}
