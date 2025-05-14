package main

import (
	"fmt"
	"math"
)

type Minheap struct {
	Length int
	data   []int
}

func (m *Minheap) insert(value int) {
	m.data = append(m.data, value)
	m.heapifyUp(m.Length)
	m.Length++
}

func (m *Minheap) delete() int {
	if m.Length == 0 {
		return -1
	}
	out := m.data[0]
	m.Length--
	if m.Length == 0 {
		m.data = []int{}
		return out
	}
	m.data[0] = m.data[m.Length]
	m.data = m.data[:m.Length]
	m.heapifyDown(0)
	return out
}

func (m *Minheap) heapifyDown(idx int) {
	if idx >= m.Length {
		return
	}

	lIdx := m.leftChild(idx)
	rIdx := m.rightChild(idx)
	smallest := idx

	if lIdx < m.Length && m.data[lIdx] < m.data[smallest] {
		smallest = lIdx
	}
	if rIdx < m.Length && m.data[rIdx] < m.data[smallest] {
		smallest = rIdx
	}

	if smallest != idx {
		m.data[idx], m.data[smallest] = m.data[smallest], m.data[idx]
		m.heapifyDown(smallest)
	}
}

func (m *Minheap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	p := m.parent(idx)
	if m.data[p] > m.data[idx] {
		m.data[p], m.data[idx] = m.data[idx], m.data[p]
		m.heapifyUp(p)
	}
}

func (m *Minheap) parent(idx int) int {
	return int(math.Floor(float64(idx-1) / 2))
}

func (m *Minheap) leftChild(idx int) int {
	return idx*2 + 1
}

func (m *Minheap) rightChild(idx int) int {
	return idx*2 + 2
}

func main() {
	heap := &Minheap{}

	heap.insert(5)
	heap.insert(3)
	heap.insert(8)
	heap.insert(1)

	for heap.Length > 0 {
		fmt.Println(heap.delete())
	}
}
