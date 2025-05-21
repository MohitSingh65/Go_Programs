package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LRU[K comparable, V any] struct {
	length int
}

func (c *LRU[K, V]) update(key K, value V) {
	fmt.Print("Hello")
}

func main() {
	fmt.Println("Hello, world!")
}
