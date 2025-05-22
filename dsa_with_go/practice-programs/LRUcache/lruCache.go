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
	capacity      int
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
	node := c.lookup[key]
	if node == nil {
		node = createNode(value)
		c.length++
		c.prepend(node)
		c.trimCache()
		c.lookup[key] = node
		c.reverseLookup[node] = &key
		node.value = value
	} else {
		c.detach(node)
		c.prepend(node)
	}
}

func (c *LRU[K, V]) get(key K) (V, bool) {
	node, ok := c.lookup[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.detach(node)
	c.prepend(node)

	return node.value, true
}

func (c *LRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if c.length == 1 {
		c.tail = c.head
		c.head = nil
	}

	if c.head == node {
		c.head = c.head.next
	}

	if c.tail == node {
		c.tail = c.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (c *LRU[K, V]) prepend(node *Node[V]) {
	if c.head == nil {
		c.head = c.tail
		c.tail = node
		return
	}

	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRU[K, V]) trimCache() {
	if c.length <= c.capacity {
		return
	}

	tail := c.tail
	c.detach(c.tail)

	key, ok := c.reverseLookup[tail]
	if !ok {
		return
	}
	delete(c.lookup, *key)
	delete(c.reverseLookup, tail)
	c.length--
}
