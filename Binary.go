package main

import (
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// InOrderTraversal performs in-order traversal of the binary tree
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left)
	fmt.Print(root.Value, " ")
	InOrderTraversal(root.Right)
}

func main() {
	// Creating a simple binary tree
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	// Performing in-order traversal
	fmt.Println("In-Order Traversal:")
	InOrderTraversal(root)
}
