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

// Queue interface defines the methods a queue must implement
type Queue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, error)
	Front() (interface{}, error)
	IsEmpty() bool
	Size() int
}

// SliceQueue implements a queue using a slice
type SliceQueue struct {
	items []interface{}
}

// NewSliceQueue creates a new queue using slice implementation
func NewSliceQueue() *SliceQueue {
	return &SliceQueue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an item to the end of the queue
func (q *SliceQueue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue
func (q *SliceQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front returns the front item without removing it
func (q *SliceQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.items[0], nil
}

// IsEmpty checks if the queue is empty
func (q *SliceQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *SliceQueue) Size() int {
	return len(q.items)
}

// QueueNode represents a node in the linked list queue
type QueueNode struct {
	data interface{}
	next *QueueNode
}

// LinkedQueue implements a queue using a linked list
type LinkedQueue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

// NewLinkedQueue creates a new queue using linked list implementation
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// Enqueue adds an item to the end of the queue
func (q *LinkedQueue) Enqueue(item interface{}) {
	newNode := &QueueNode{
		data: item,
		next: nil,
	}

	if q.IsEmpty() {
		q.front = newNode
	} else {
		q.rear.next = newNode
	}

	q.rear = newNode
	q.size++
}

// Dequeue removes and returns the front item from the queue
func (q *LinkedQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	item := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return item, nil
}

// Front returns the front item without removing it
func (q *LinkedQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.front.data, nil
}

// IsEmpty checks if the queue is empty
func (q *LinkedQueue) IsEmpty() bool {
	return q.front == nil
}

// Size returns the number of items in the queue
func (q *LinkedQueue) Size() int {
	return q.size
}

// NewQueue creates a new queue (using the slice implementation by default)
func NewQueue() Queue {
	return NewSliceQueue()
}
