/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedListOperations(t *testing.T) {
	// Create a new linked list
	list := New()

	// Initial state
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	if list.Length() != 0 {
		t.Errorf("Expected length 0, got %d", list.Length())
	}

	// Insert at beginning
	list.InsertAtBeginning(10)

	if list.IsEmpty() {
		t.Error("List should not be empty after insertion")
	}

	if list.Length() != 1 {
		t.Errorf("Expected length 1, got %d", list.Length())
	}

	if list.Head.Data != 10 {
		t.Errorf("Expected head data 10, got %v", list.Head.Data)
	}

	// Insert at end
	list.InsertAtEnd(20)

	if list.Length() != 2 {
		t.Errorf("Expected length 2, got %d", list.Length())
	}

	tail := list.GetTail()
	if tail.Data != 20 {
		t.Errorf("Expected tail data 20, got %v", tail.Data)
	}

	// Insert after
	success := list.InsertAfter(10, 15)
	if !success {
		t.Error("InsertAfter should return true for successful insertion")
	}

	if list.Length() != 3 {
		t.Errorf("Expected length 3, got %d", list.Length())
	}

	// Check the list structure: 10 -> 15 -> 20
	expected := []interface{}{10, 15, 20}
	actual := list.ToSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected list %v, got %v", expected, actual)
	}

	// Search for existing element
	node := list.Search(15)
	if node == nil {
		t.Error("Search should find element 15")
	} else if node.Data != 15 {
		t.Errorf("Search returned wrong node, expected 15 got %v", node.Data)
	}

	// Search for non-existing element
	node = list.Search(100)
	if node != nil {
		t.Error("Search should not find element 100")
	}

	// Delete middle node
	success = list.DeleteNode(15)
	if !success {
		t.Error("DeleteNode should return true for successful deletion")
	}

	if list.Length() != 2 {
		t.Errorf("Expected length 2 after deletion, got %d", list.Length())
	}

	expected = []interface{}{10, 20}
	actual = list.ToSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected list %v after deletion, got %v", expected, actual)
	}

	// Delete head
	success = list.DeleteNode(10)
	if !success {
		t.Error("DeleteNode should return true for successful deletion of head")
	}

	if list.Length() != 1 {
		t.Errorf("Expected length 1 after deletion, got %d", list.Length())
	}

	if list.Head.Data != 20 {
		t.Errorf("Expected head data 20 after deletion, got %v", list.Head.Data)
	}

	// Delete non-existing node
	success = list.DeleteNode(100)
	if success {
		t.Error("DeleteNode should return false for non-existing element")
	}

	// Delete last node
	success = list.DeleteNode(20)
	if !success {
		t.Error("DeleteNode should return true for successful deletion of last node")
	}

	if !list.IsEmpty() {
		t.Error("List should be empty after deleting all elements")
	}

	if list.Length() != 0 {
		t.Errorf("Expected length 0 after deletion, got %d", list.Length())
	}
}

func TestReverse(t *testing.T) {
	// Create a list with elements
	list := FromSlice([]interface{}{1, 2, 3, 4, 5})

	// Reverse the list
	list.Reverse()

	// Check the result
	expected := []interface{}{5, 4, 3, 2, 1}
	actual := list.ToSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected reversed list %v, got %v", expected, actual)
	}

	// Test with an empty list
	emptyList := New()
	emptyList.Reverse() // Should not crash

	if !emptyList.IsEmpty() {
		t.Error("Reversed empty list should still be empty")
	}

	// Test with a single element
	singleList := FromSlice([]interface{}{42})
	singleList.Reverse()

	expected = []interface{}{42}
	actual = singleList.ToSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected single element list to remain %v, got %v", expected, actual)
	}
}

func TestFromSlice(t *testing.T) {
	// Create a list from a slice
	slice := []interface{}{1, 2, 3, 4, 5}
	list := FromSlice(slice)

	// Check the result
	actual := list.ToSlice()

	if !reflect.DeepEqual(slice, actual) {
		t.Errorf("Expected list from slice %v, got %v", slice, actual)
	}

	// Test with an empty slice
	emptySlice := []interface{}{}
	emptyList := FromSlice(emptySlice)

	if !emptyList.IsEmpty() {
		t.Error("List from empty slice should be empty")
	}
}
