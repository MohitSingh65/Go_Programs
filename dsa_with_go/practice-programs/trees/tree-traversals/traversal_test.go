package main

import (
	"reflect"
	"testing"
)

func buildSampleTree() *Node {
	// Construct the following tree:
	//         1
	//        / \
	//       2   3
	//      / \
	//     4   5
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	return root
}

func TestPreOrderTraversal(t *testing.T) {
	root := buildSampleTree()
	expected := []int{1, 2, 4, 5, 3}
	result := PreOrderTraversal(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderTraversal failed. Expected %v, got %v", expected, result)
	}
}

func TestInOrderTraversal(t *testing.T) {
	root := buildSampleTree()
	expected := []int{4, 2, 5, 1, 3}
	result := InOrderTraversal(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrderTraversal failed. Expected %v, got %v", expected, result)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	root := buildSampleTree()
	expected := []int{4, 5, 2, 3, 1}
	result := PostOrderTraversal(root)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostOrderTraversal failed. Expected %v, got %v", expected, result)
	}
}
