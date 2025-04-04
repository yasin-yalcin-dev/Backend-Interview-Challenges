# Binary Search Tree

## Problem
Implement a Binary Search Tree (BST) data structure with standard operations like insertion, deletion, search, and traversal. A Binary Search Tree is a node-based binary tree data structure that has the following properties:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees are also binary search trees.

## Requirements
1. Implement a BST node structure containing data and pointers to left and right children
2. Implement a BST structure with a root node
3. Support the following operations:
   - Insert - Add a new node with a given value
   - Delete - Remove a node with a given value
   - Search - Find a node with a given value
   - Min/Max - Find the minimum/maximum value in the tree
   - InOrderTraversal - Visit all nodes in-order (left, root, right)
   - PreOrderTraversal - Visit all nodes pre-order (root, left, right)
   - PostOrderTraversal - Visit all nodes post-order (left, right, root)
   - LevelOrderTraversal - Visit all nodes level by level
   - Height - Calculate the height of the tree
   - IsEmpty - Check if the tree is empty
   - Size - Get the number of nodes in the tree

## Examples
```go
// Create a new BST
tree := NewBST()

// Insert elements
tree.Insert(50)
tree.Insert(30)
tree.Insert(70)
tree.Insert(20)
tree.Insert(40)
tree.Insert(60)
tree.Insert(80)

// Search for an element
found := tree.Search(40)  // Returns true
found = tree.Search(90)   // Returns false

// Find min and max
min := tree.Min()  // Returns 20
max := tree.Max()  // Returns 80

// Print in different traversal orders
fmt.Println("In-order traversal:")
tree.InOrderTraversal()  // Output: 20 30 40 50 60 70 80

// Delete an element
tree.Delete(30)

// Print again to see the change
fmt.Println("In-order traversal after deletion:")
tree.InOrderTraversal()  // Output: 20 40 50 60 70 80