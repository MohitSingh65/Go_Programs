package main

import "testing"

func buildTestTree() *Node {
	// Tree structure:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	return root
}

func TestBfs(t *testing.T) {
	root := buildTestTree()

	tests := []struct {
		needle   int
		expected bool
	}{
		{needle: 1, expected: true},
		{needle: 3, expected: true},
		{needle: 5, expected: true},
		{needle: 42, expected: false},
	}

	for _, test := range tests {
		got := Bfs(root, test.needle)
		if got != test.expected {
			t.Errorf("Bfs(%d) = %v; want %v", test.needle, got, test.expected)
		}
	}
}
