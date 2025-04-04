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
	"reflect"
	"testing"
)

func TestBSTBasicOperations(t *testing.T) {
	// Create a new BST
	bst := NewBST()

	// Test initial state
	if !bst.IsEmpty() {
		t.Error("New BST should be empty")
	}

	if bst.Size() != 0 {
		t.Errorf("Expected size 0, got %d", bst.Size())
	}

	// Test insert
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(val)
	}

	if bst.IsEmpty() {
		t.Error("BST should not be empty after insertion")
	}

	if bst.Size() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), bst.Size())
	}

	// Test search
	for _, val := range values {
		if !bst.Search(val) {
			t.Errorf("Search should find element %d", val)
		}
	}

	if bst.Search(100) {
		t.Error("Search should not find element 100")
	}

	// Test min and max
	min, err := bst.Min()
	if err != nil {
		t.Errorf("Min should not return error: %v", err)
	}
	if min != 20 {
		t.Errorf("Expected min 20, got %d", min)
	}

	max, err := bst.Max()
	if err != nil {
		t.Errorf("Max should not return error: %v", err)
	}
	if max != 80 {
		t.Errorf("Expected max 80, got %d", max)
	}

	// Test height
	if bst.Height() != 3 {
		t.Errorf("Expected height 3, got %d", bst.Height())
	}

	// Test is valid BST
	if !bst.IsBST() {
		t.Error("Tree should be a valid BST")
	}
}

func TestBSTTraversals(t *testing.T) {
	// Create a BST with a known structure
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(val)
	}

	// Test in-order traversal (should be sorted)
	expectedInOrder := []int{20, 30, 40, 50, 60, 70, 80}
	inOrder := bst.InOrderTraversal()

	if !reflect.DeepEqual(inOrder, expectedInOrder) {
		t.Errorf("In-order traversal mismatch.\nExpected: %v\nGot: %v", expectedInOrder, inOrder)
	}

	// Test pre-order traversal
	expectedPreOrder := []int{50, 30, 20, 40, 70, 60, 80}
	preOrder := bst.PreOrderTraversal()

	if !reflect.DeepEqual(preOrder, expectedPreOrder) {
		t.Errorf("Pre-order traversal mismatch.\nExpected: %v\nGot: %v", expectedPreOrder, preOrder)
	}

	// Test post-order traversal
	expectedPostOrder := []int{20, 40, 30, 60, 80, 70, 50}
	postOrder := bst.PostOrderTraversal()

	if !reflect.DeepEqual(postOrder, expectedPostOrder) {
		t.Errorf("Post-order traversal mismatch.\nExpected: %v\nGot: %v", expectedPostOrder, postOrder)
	}

	// Test level-order traversal
	expectedLevelOrder := []int{50, 30, 70, 20, 40, 60, 80}
	levelOrder := bst.LevelOrderTraversal()

	if !reflect.DeepEqual(levelOrder, expectedLevelOrder) {
		t.Errorf("Level-order traversal mismatch.\nExpected: %v\nGot: %v", expectedLevelOrder, levelOrder)
	}
}

func TestBSTDeletion(t *testing.T) {
	// Create a BST
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		bst.Insert(val)
	}

	// Case 1: Delete a leaf node
	result := bst.Delete(20)
	if !result {
		t.Error("Delete should return true for existing element")
	}

	if bst.Search(20) {
		t.Error("Element 20 should be deleted")
	}

	if bst.Size() != 6 {
		t.Errorf("Expected size 6 after deletion, got %d", bst.Size())
	}

	expectedAfterLeafDeletion := []int{30, 40, 50, 60, 70, 80}
	if !reflect.DeepEqual(bst.InOrderTraversal(), expectedAfterLeafDeletion) {
		t.Errorf("Incorrect tree after leaf deletion.\nExpected: %v\nGot: %v",
			expectedAfterLeafDeletion, bst.InOrderTraversal())
	}

	// Case 2: Delete a node with one child
	// First add a child to node 60
	bst.Insert(65)
	// Then delete 60
	result = bst.Delete(60)
	if !result {
		t.Error("Delete should return true for existing element")
	}

	if bst.Search(60) {
		t.Error("Element 60 should be deleted")
	}

	expectedAfterOneChildDeletion := []int{30, 40, 50, 65, 70, 80}
	if !reflect.DeepEqual(bst.InOrderTraversal(), expectedAfterOneChildDeletion) {
		t.Errorf("Incorrect tree after one-child deletion.\nExpected: %v\nGot: %v",
			expectedAfterOneChildDeletion, bst.InOrderTraversal())
	}

	// Case 3: Delete a node with two children
	result = bst.Delete(70)
	if !result {
		t.Error("Delete should return true for existing element")
	}

	if bst.Search(70) {
		t.Error("Element 70 should be deleted")
	}

	expectedAfterTwoChildrenDeletion := []int{30, 40, 50, 65, 80}
	if !reflect.DeepEqual(bst.InOrderTraversal(), expectedAfterTwoChildrenDeletion) {
		t.Errorf("Incorrect tree after two-children deletion.\nExpected: %v\nGot: %v",
			expectedAfterTwoChildrenDeletion, bst.InOrderTraversal())
	}

	// Case 4: Delete the root
	result = bst.Delete(50)
	if !result {
		t.Error("Delete should return true for existing element")
	}

	if bst.Search(50) {
		t.Error("Root element 50 should be deleted")
	}

	expectedAfterRootDeletion := []int{30, 40, 65, 80}
	if !reflect.DeepEqual(bst.InOrderTraversal(), expectedAfterRootDeletion) {
		t.Errorf("Incorrect tree after root deletion.\nExpected: %v\nGot: %v",
			expectedAfterRootDeletion, bst.InOrderTraversal())
	}

	// Try to delete non-existing element
	result = bst.Delete(100)
	if result {
		t.Error("Delete should return false for non-existing element")
	}
}

func TestEmptyBST(t *testing.T) {
	// Create an empty BST
	bst := NewBST()

	// Test operations on empty tree
	_, err := bst.Min()
	if err == nil {
		t.Error("Min should return error on empty tree")
	}

	_, err = bst.Max()
	if err == nil {
		t.Error("Max should return error on empty tree")
	}

	if bst.Height() != 0 {
		t.Errorf("Height of empty tree should be 0, got %d", bst.Height())
	}

	if !bst.IsBST() {
		t.Error("Empty tree should be a valid BST")
	}

	// Test traversals on empty tree
	if len(bst.InOrderTraversal()) != 0 {
		t.Error("In-order traversal of empty tree should return empty slice")
	}

	if len(bst.PreOrderTraversal()) != 0 {
		t.Error("Pre-order traversal of empty tree should return empty slice")
	}

	if len(bst.PostOrderTraversal()) != 0 {
		t.Error("Post-order traversal of empty tree should return empty slice")
	}

	if len(bst.LevelOrderTraversal()) != 0 {
		t.Error("Level-order traversal of empty tree should return empty slice")
	}

	// Test deletion on empty tree
	if bst.Delete(10) {
		t.Error("Delete on empty tree should return false")
	}
}
