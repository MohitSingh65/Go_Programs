package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Node{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (t *Trie) Autocomplete(prefix string) []string {
	curr := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return nil
		}
		curr = curr.children[idx]
	}

	var result []string
	var dfs func(node *Node, path string)
	dfs = func(node *Node, path string) {
		if node.isEnd {
			result = append(result, path)
		}

		for i, child := range node.children {
			if child != nil {
				dfs(child, path+string('a'+i))
			}
		}
	}

	dfs(curr, prefix)
	return result
}

func main() {
	trie := NewTrie()
	words := []string{"go", "gone", "god", "good", "gum"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Autocomplete("go"))
	fmt.Println(trie.Autocomplete("gu"))
	fmt.Println(trie.Autocomplete("z"))
}
