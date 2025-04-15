package main

import "fmt"

const maxSize = 100

type Stack struct {
	arr [maxSize]int
	top int
}

func (s *Stack) underflow() bool {
	return s.top == 0
}

func (s *Stack) overflow() bool {
	return s.top == maxSize
}

func (s *Stack) push(value int) {
	if s.overflow() {
		fmt.Println("Stack is full!")
	}
	s.arr[s.top] = value
	s.top++
}

func (s *Stack) pop() int {
	if s.underflow() {
		fmt.Println("Stack is empty!")
	}
	s.top--
	return s.arr[s.top]
}

func (s *Stack) peek() int {
	if s.underflow() {
		fmt.Println("Stack is empty!")
	}
	return s.arr[s.top-1]
}

func main() {
	s := Stack{}
	s.push(10)
	s.push(20)
	s.push(30)

	fmt.Printf("Top element: %v\n", s.peek())
	fmt.Printf("Stack size: %v\n", s.top)

	s.pop()
	fmt.Println("Popped!")
	s.pop()
	fmt.Println("Popped!")
	s.push(40)

	fmt.Printf("Top element: %v\n", s.peek())
	fmt.Printf("Stack size: %v\n", s.top)

}
