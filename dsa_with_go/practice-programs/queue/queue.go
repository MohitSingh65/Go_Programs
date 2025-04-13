package main

import (
	"fmt"
)

const maxSize = 100

type Queue struct {
	arr   [maxSize]int
	count int
	front int
	rear  int
}

func (q *Queue) isEmpty() bool {
	return q.count == 0
}

func (q *Queue) isFull() bool {
	return q.count == maxSize
}

func (q *Queue) Enqueue(value int) {
	if q.isFull() {
		fmt.Println("Queue overflow!")
	}
	q.arr[q.rear] = value
	q.rear = (q.rear + 1) % maxSize
	q.count++
}

func (q *Queue) Dequeue() (value int) {
	if q.isEmpty() {
		fmt.Println("Queue underflow!")
		return -1
	}
	value = q.arr[q.front]
	q.front = (q.front + 1) % maxSize
	q.count--
	return value
}

func (q *Queue) Peek() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty!")
		return -1
	}
	return q.arr[q.front]
}

func (q *Queue) Size() int {
	return q.count
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Printf("Front element: %v\n", q.Peek())
	fmt.Printf("Queue size: %v\n", q.Size())

	q.Dequeue()
	fmt.Println("Dequeued")
	q.Dequeue()
	fmt.Println("Dequeued")

	q.Enqueue(40)
	fmt.Printf("Front element: %v\n", q.Peek())
	fmt.Printf("Queue size: %v\n", q.Size())

}
