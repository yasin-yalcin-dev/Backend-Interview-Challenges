# Stack & Queue

## Problem
Implement two fundamental data structures: Stack (LIFO - Last In, First Out) and Queue (FIFO - First In, First Out) with their standard operations.

## Requirements
1. **Stack Implementation:**
   - Push - Add an element to the top of the stack
   - Pop - Remove the top element from the stack
   - Peek - Get the top element without removing it
   - IsEmpty - Check if the stack is empty
   - Size - Get the number of elements in the stack

2. **Queue Implementation:**
   - Enqueue - Add an element to the end of the queue
   - Dequeue - Remove the element from the front of the queue
   - Front - Get the front element without removing it
   - IsEmpty - Check if the queue is empty
   - Size - Get the number of elements in the queue

## Examples

**Stack Example:**
```go
// Create a new stack
stack := NewStack()

// Push elements
stack.Push(1)
stack.Push(2)
stack.Push(3)

// Peek at the top element
fmt.Println(stack.Peek())  // Output: 3

// Pop elements
fmt.Println(stack.Pop())  // Output: 3
fmt.Println(stack.Pop())  // Output: 2
fmt.Println(stack.Pop())  // Output: 1

// Check if empty
fmt.Println(stack.IsEmpty())  // Output: true
Queue Example:
goCopy// Create a new queue
queue := NewQueue()

// Enqueue elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Get the front element
fmt.Println(queue.Front())  // Output: 1

// Dequeue elements
fmt.Println(queue.Dequeue())  // Output: 1
fmt.Println(queue.Dequeue())  // Output: 2
fmt.Println(queue.Dequeue())  // Output: 3

// Check if empty
fmt.Println(queue.IsEmpty())  // Output: true