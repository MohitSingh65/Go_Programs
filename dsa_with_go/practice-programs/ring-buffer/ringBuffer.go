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
}
