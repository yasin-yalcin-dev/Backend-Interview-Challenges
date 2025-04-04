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
	"fmt"
	"strings"
)

// Node represents a node in a linked list
type Node struct {
	Data interface{}
	Next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
	size int
}

// New creates a new empty linked list
func New() *LinkedList {
	return &LinkedList{
		Head: nil,
		size: 0,
	}
}

// IsEmpty checks if the linked list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

// Length returns the number of nodes in the linked list
func (ll *LinkedList) Length() int {
	return ll.size
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{
		Data: data,
		Next: ll.Head,
	}
	ll.Head = newNode
	ll.size++
}

// InsertAtEnd inserts a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if ll.IsEmpty() {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.size++
}

// InsertAfter inserts a new node after the node containing the given value
// Returns true if insertion was successful, false if the value was not found
func (ll *LinkedList) InsertAfter(value, data interface{}) bool {
	if ll.IsEmpty() {
		return false
	}

	current := ll.Head
	for current != nil {
		if current.Data == value {
			newNode := &Node{
				Data: data,
				Next: current.Next,
			}
			current.Next = newNode
			ll.size++
			return true
		}
		current = current.Next
	}
	return false
}

// DeleteNode deletes the first node with the given value
// Returns true if deletion was successful, false if the value was not found
func (ll *LinkedList) DeleteNode(value interface{}) bool {
	if ll.IsEmpty() {
		return false
	}

	// Special case: delete head
	if ll.Head.Data == value {
		ll.Head = ll.Head.Next
		ll.size--
		return true
	}

	// General case: delete node after x
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == value {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}
	return false
}

// Search finds the first node containing the given value
// Returns nil if the value is not found
func (ll *LinkedList) Search(value interface{}) *Node {
	if ll.IsEmpty() {
		return nil
	}

	current := ll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// GetHead returns the head node of the list
func (ll *LinkedList) GetHead() *Node {
	return ll.Head
}

// GetTail returns the tail node of the list
// Returns nil if the list is empty
func (ll *LinkedList) GetTail() *Node {
	if ll.IsEmpty() {
		return nil
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// Display prints all elements in the list
func (ll *LinkedList) Display() {
	if ll.IsEmpty() {
		fmt.Println("List is empty")
		return
	}

	current := ll.Head
	elements := []string{}

	for current != nil {
		elements = append(elements, fmt.Sprintf("%v", current.Data))
		current = current.Next
	}

	fmt.Println(strings.Join(elements, " -> "))
}

// ToSlice converts the linked list to a slice
func (ll *LinkedList) ToSlice() []interface{} {
	result := make([]interface{}, 0, ll.size)
	current := ll.Head

	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}

	return result
}

// FromSlice creates a linked list from a slice
func FromSlice(items []interface{}) *LinkedList {
	list := New()
	for _, item := range items {
		list.InsertAtEnd(item)
	}
	return list
}

// Reverse reverses the linked list in place
func (ll *LinkedList) Reverse() {
	if ll.IsEmpty() || ll.Head.Next == nil {
		return
	}

	var prev *Node
	current := ll.Head
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}
