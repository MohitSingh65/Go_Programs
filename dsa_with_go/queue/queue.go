package main

import "fmt"

// Queue structure using a slice
type Queue struct {
	items []int
}

// Enqueue adds an element at the back of the queue
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:] // Remove the first element
	return item, nil
}

// Peek returns the front element without removing it
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Queue:", q.items)

	item, _ := q.Dequeue()
	fmt.Println("Dequeued:", item)
	fmt.Println("Queue after dequeue:", q.items)

	front, _ := q.Peek()
	fmt.Println("Front element:", front)

	fmt.Println("Is queue empty?", q.IsEmpty())
	fmt.Println("Queue size:", q.Size())
}
