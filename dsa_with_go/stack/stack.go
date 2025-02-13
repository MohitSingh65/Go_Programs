package main

import "fmt"

type Stack struct {
	items []int
}

// Push adds an element to the top
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

// Peek returns the top element without removing items
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Top element:", stack.Peek())
	fmt.Println("Popped:", stack.Pop())
	fmt.Println("Popped:", stack.Pop())
	fmt.Println("Is Empty?", stack.IsEmpty())
}
