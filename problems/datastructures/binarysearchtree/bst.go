/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package binarysearchtree

import (
	"fmt"
	"math"
)

// Node represents a node in a binary search tree
type Node struct {
	Value       int
	Left, Right *Node
}

// BST represents a binary search tree
type BST struct {
	Root *Node
	size int
}

// NewBST creates a new empty binary search tree
func NewBST() *BST {
	return &BST{
		Root: nil,
		size: 0,
	}
}

// IsEmpty checks if the tree is empty
func (bst *BST) IsEmpty() bool {
	return bst.Root == nil
}

// Size returns the number of nodes in the tree
func (bst *BST) Size() int {
	return bst.size
}

// Insert adds a new value to the tree
func (bst *BST) Insert(value int) {
	newNode := &Node{Value: value}

	// If tree is empty, set root
	if bst.IsEmpty() {
		bst.Root = newNode
		bst.size++
		return
	}

	// Otherwise, find the correct spot
	insertNode(bst.Root, newNode)
	bst.size++
}

// insertNode is a helper recursive function to insert a node in the tree
func insertNode(node, newNode *Node) {
	// If value is less than the current node, go left
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		// If value is greater or equal, go right
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

// Search checks if a value exists in the tree
func (bst *BST) Search(value int) bool {
	return searchNode(bst.Root, value)
}

// searchNode is a helper recursive function to search for a value
func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return searchNode(node.Left, value)
	}

	return searchNode(node.Right, value)
}

// Min finds the minimum value in the tree
func (bst *BST) Min() (int, error) {
	if bst.IsEmpty() {
		return 0, fmt.Errorf("tree is empty")
	}

	return findMin(bst.Root).Value, nil
}

// findMin is a helper function to find the leftmost node
func findMin(node *Node) *Node {
	current := node

	// Keep going left until we reach a leaf
	for current.Left != nil {
		current = current.Left
	}

	return current
}

// Max finds the maximum value in the tree
func (bst *BST) Max() (int, error) {
	if bst.IsEmpty() {
		return 0, fmt.Errorf("tree is empty")
	}

	return findMax(bst.Root).Value, nil
}

// findMax is a helper function to find the rightmost node
func findMax(node *Node) *Node {
	current := node

	// Keep going right until we reach a leaf
	for current.Right != nil {
		current = current.Right
	}

	return current
}

// Delete removes a value from the tree
func (bst *BST) Delete(value int) bool {
	if bst.IsEmpty() {
		return false
	}

	// Keep track if we actually deleted something
	found := false
	bst.Root, found = deleteNode(bst.Root, value)

	if found {
		bst.size--
	}

	return found
}

// deleteNode is a helper recursive function to delete a node
func deleteNode(node *Node, value int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	var found bool

	if value < node.Value {
		// Value is in the left subtree
		node.Left, found = deleteNode(node.Left, value)
	} else if value > node.Value {
		// Value is in the right subtree
		node.Right, found = deleteNode(node.Right, value)
	} else {
		// We found the node to delete
		found = true

		// Case 1: Leaf node (no children)
		if node.Left == nil && node.Right == nil {
			return nil, found
		}

		// Case 2: Node with only one child
		if node.Left == nil {
			return node.Right, found
		}

		if node.Right == nil {
			return node.Left, found
		}

		// Case 3: Node with two children
		// Find the inorder successor (smallest node in right subtree)
		minRight := findMin(node.Right)

		// Replace this node's value with successor's value
		node.Value = minRight.Value

		// Delete the successor
		node.Right, _ = deleteNode(node.Right, minRight.Value)
	}

	return node, found
}

// InOrderTraversal visits all nodes in-order (left, root, right)
func (bst *BST) InOrderTraversal() []int {
	result := make([]int, 0, bst.size)
	inOrderTraversal(bst.Root, &result)
	return result
}

// inOrderTraversal is a helper recursive function for in-order traversal
func inOrderTraversal(node *Node, result *[]int) {
	if node != nil {
		inOrderTraversal(node.Left, result)
		*result = append(*result, node.Value)
		inOrderTraversal(node.Right, result)
	}
}

// PreOrderTraversal visits all nodes pre-order (root, left, right)
func (bst *BST) PreOrderTraversal() []int {
	result := make([]int, 0, bst.size)
	preOrderTraversal(bst.Root, &result)
	return result
}

// preOrderTraversal is a helper recursive function for pre-order traversal
func preOrderTraversal(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		preOrderTraversal(node.Left, result)
		preOrderTraversal(node.Right, result)
	}
}

// PostOrderTraversal visits all nodes post-order (left, right, root)
func (bst *BST) PostOrderTraversal() []int {
	result := make([]int, 0, bst.size)
	postOrderTraversal(bst.Root, &result)
	return result
}

// postOrderTraversal is a helper recursive function for post-order traversal
func postOrderTraversal(node *Node, result *[]int) {
	if node != nil {
		postOrderTraversal(node.Left, result)
		postOrderTraversal(node.Right, result)
		*result = append(*result, node.Value)
	}
}

// LevelOrderTraversal visits all nodes level by level
func (bst *BST) LevelOrderTraversal() []int {
	result := make([]int, 0, bst.size)

	if bst.IsEmpty() {
		return result
	}

	// Create a queue for level order traversal
	queue := make([]*Node, 0)
	queue = append(queue, bst.Root)

	for len(queue) > 0 {
		// Dequeue a node
		node := queue[0]
		queue = queue[1:]

		// Add the node's value to the result
		result = append(result, node.Value)

		// Enqueue left child
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// Enqueue right child
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// Height returns the height of the tree
func (bst *BST) Height() int {
	if bst.IsEmpty() {
		return 0
	}

	return calculateHeight(bst.Root)
}

// calculateHeight is a helper recursive function to find the height
func calculateHeight(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := calculateHeight(node.Left)
	rightHeight := calculateHeight(node.Right)

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// IsBST checks if the tree is a valid binary search tree
func (bst *BST) IsBST() bool {
	return isBSTUtil(bst.Root, math.MinInt, math.MaxInt)
}

// isBSTUtil is a helper recursive function to check if a tree is a BST
func isBSTUtil(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	// Check if the current node's value is within the allowed range
	if node.Value < min || node.Value > max {
		return false
	}

	// Check the left and right subtrees
	return isBSTUtil(node.Left, min, node.Value-1) &&
		isBSTUtil(node.Right, node.Value+1, max)
}

// PrintTree prints a visual representation of the tree
func (bst *BST) PrintTree() {
	printNode(bst.Root, 0)
}

// printNode is a helper function to print a tree visually
func printNode(node *Node, level int) {
	if node == nil {
		return
	}

	// Print right branch first (so it appears at the top)
	printNode(node.Right, level+1)

	// Print the current node
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Println(node.Value)

	// Print left branch
	printNode(node.Left, level+1)
}
