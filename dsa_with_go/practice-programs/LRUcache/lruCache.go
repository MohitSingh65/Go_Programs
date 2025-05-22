package main

type Node[V any] struct {
	value V
	next  *Node[V]
	prev  *Node[V]
}

func createNode[V any](value V) *Node[V] {
	return &Node[V]{value: value}
}

type LRU[K comparable, V any] struct {
	length        int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]*K
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		length:        0,
		head:          nil,
		tail:          nil,
		lookup:        make(map[K]*Node[V]),
		reverseLookup: make(map[*Node[V]]*K),
	}
}

func (c *LRU[K, V]) update(key K, value V) {
}

func (c *LRU[K, V]) get(key K) (V, bool) {
	node, ok := c.lookup[key]
	if !ok {
		var zero V
		return zero, false
	}
	return node.value, true
}
