/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package stackqueue

import "errors"

// Stack interface defines the methods a stack must implement
type Stack interface {
	Push(item interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
}

// SliceStack implements a stack using a slice
type SliceStack struct {
	items []interface{}
}

// NewSliceStack creates a new stack using slice implementation
func NewSliceStack() *SliceStack {
	return &SliceStack{
		items: make([]interface{}, 0),
	}
}

// Push adds an item to the top of the stack
func (s *SliceStack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *SliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the top item without removing it
func (s *SliceStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *SliceStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *SliceStack) Size() int {
	return len(s.items)
}

// ListNode represents a node in the linked list stack
type ListNode struct {
	data interface{}
	next *ListNode
}

// LinkedStack implements a stack using a linked list
type LinkedStack struct {
	top  *ListNode
	size int
}

// NewLinkedStack creates a new stack using linked list implementation
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top:  nil,
		size: 0,
	}
}

// Push adds an item to the top of the stack
func (s *LinkedStack) Push(item interface{}) {
	newNode := &ListNode{
		data: item,
		next: s.top,
	}
	s.top = newNode
	s.size++
}

// Pop removes and returns the top item from the stack
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	item := s.top.data
	s.top = s.top.next
	s.size--
	return item, nil
}

// Peek returns the top item without removing it
func (s *LinkedStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.top.data, nil
}

// IsEmpty checks if the stack is empty
func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of items in the stack
func (s *LinkedStack) Size() int {
	return s.size
}

// NewStack creates a new stack (using the slice implementation by default)
func NewStack() Stack {
	return NewSliceStack()
}
