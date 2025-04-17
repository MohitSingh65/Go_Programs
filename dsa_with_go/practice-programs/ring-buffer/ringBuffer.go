package main

import (
	"errors"
)

type RingBuffer struct {
	buffer []interface{}
	size   int
	head   int
	tail   int
	isFull bool
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]interface{}, size),
		size:   size,
	}
}

func (r *RingBuffer) Enqueue(item interface{}) error {
	if r.isFull {
		return errors.New("buffer is null")
	}

	r.buffer[r.head] = item
	r.head = (r.head + 1) % r.size

	if r.head == r.tail {
		r.isFull = true
	}
	return nil
}

func (r *RingBuffer) Dequeue() (interface{}, error) {
	if r.IsEmpty {
		return nil, errors.New("buffer is empty")
	}

	item := r.buffer[r.tail]
	r.tail = (r.tail + 1) % r.size
	r.isFull = false

	return item, nil
}

func (r *RingBuffer) IsEmpty() bool {
	return !r.isFull && r.head == r.tail
}

func (r *RingBuffer) isFull() bool {
	return r.IsFull
}
