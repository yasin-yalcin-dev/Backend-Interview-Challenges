/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package binarysearchtree

import "fmt"

// RunExample demonstrates the binary search tree implementation
func RunExample() {
	fmt.Println("Binary Search Tree Example:")
	fmt.Println("-------------------------")

	// Create a new binary search tree
	bst := NewBST()

	// Check if it's empty
	fmt.Printf("Is the tree empty? %v\n", bst.IsEmpty())

	// Insert elements
	fmt.Println("\nInserting elements: 50, 30, 70, 20, 40, 60, 80")
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(val)
	}

	// Print the tree structure
	fmt.Println("\nTree structure (sideways):")
	bst.PrintTree()

	// Get tree information
	fmt.Printf("\nTree size: %d\n", bst.Size())
	fmt.Printf("Tree height: %d\n", bst.Height())
	fmt.Printf("Is a valid BST? %v\n", bst.IsBST())

	// Find min and max
	min, _ := bst.Min()
	max, _ := bst.Max()
	fmt.Printf("\nMinimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)

	// Search for elements
	fmt.Println("\nSearching for elements:")
	fmt.Printf("Is 40 in the tree? %v\n", bst.Search(40))
	fmt.Printf("Is 90 in the tree? %v\n", bst.Search(90))

	// Traverse the tree in different orders
	fmt.Println("\nTraversals:")
	fmt.Printf("In-order traversal:    %v\n", bst.InOrderTraversal())
	fmt.Printf("Pre-order traversal:   %v\n", bst.PreOrderTraversal())
	fmt.Printf("Post-order traversal:  %v\n", bst.PostOrderTraversal())
	fmt.Printf("Level-order traversal: %v\n", bst.LevelOrderTraversal())

	// Delete elements
	fmt.Println("\nDeleting elements:")

	// Case 1: Delete a leaf node (20)
	fmt.Println("1. Delete a leaf node (20)")
	bst.Delete(20)
	fmt.Printf("   In-order after deletion: %v\n", bst.InOrderTraversal())

	// Case 2: Delete a node with one child (60)
	fmt.Println("2. Delete a node with one child (60)")
	bst.Insert(65) // Add a child to 60
	bst.Delete(60)
	fmt.Printf("   In-order after deletion: %v\n", bst.InOrderTraversal())

	// Case 3: Delete a node with two children (30)
	fmt.Println("3. Delete a node with two children (30)")
	bst.Delete(30)
	fmt.Printf("   In-order after deletion: %v\n", bst.InOrderTraversal())

	// Case 4: Delete the root
	fmt.Println("4. Delete the root (50)")
	bst.Delete(50)
	fmt.Printf("   In-order after deletion: %v\n", bst.InOrderTraversal())

	// Print the final tree
	fmt.Println("\nFinal tree structure:")
	bst.PrintTree()

	// Practical application: Find the kth smallest element
	fmt.Println("\nPractical Application:")

	newBST := NewBST()
	for _, val := range []int{50, 30, 70, 20, 40, 60, 80, 35, 45, 55, 65} {
		newBST.Insert(val)
	}

	fmt.Printf("BST contains: %v\n", newBST.InOrderTraversal())

	k := 3
	kthSmallest := findKthSmallest(newBST, k)
	fmt.Printf("The %drd smallest element is: %d\n", k, kthSmallest)

	k = 7
	kthSmallest = findKthSmallest(newBST, k)
	fmt.Printf("The %dth smallest element is: %d\n", k, kthSmallest)
}

// findKthSmallest finds the kth smallest element in the tree
func findKthSmallest(bst *BST, k int) int {
	// Get the in-order traversal (sorted)
	inOrder := bst.InOrderTraversal()

	// Return the kth element (0-indexed)
	if k > 0 && k <= len(inOrder) {
		return inOrder[k-1]
	}

	return -1 // Invalid k
}
