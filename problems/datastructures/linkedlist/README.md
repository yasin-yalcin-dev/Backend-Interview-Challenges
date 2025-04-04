# Linked List

## Problem
Implement a singly linked list data structure with common operations like insertion, deletion, and traversal. 

## Requirements
1. Implement a Node structure containing data and a pointer to the next node
2. Implement a LinkedList structure with head pointer
3. Support the following operations:
   - InsertAtBeginning - insert a node at the beginning of the list
   - InsertAtEnd - insert a node at the end of the list
   - InsertAfter - insert a node after a given node
   - DeleteNode - delete a node with a given value
   - Search - find a node with a given value
   - Display - print all elements in the list
   - Length - return the number of nodes in the list
   - IsEmpty - check if the list is empty
   - GetHead - return the head node
   - GetTail - return the tail node

## Examples
```go
// Create a new linked list
list := linkedlist.New()

// Insert elements
list.InsertAtBeginning(1)
list.InsertAtEnd(3)
list.InsertAfter(1, 2)  // Insert 2 after 1

// Display the list
list.Display()  // Output: 1 -> 2 -> 3

// Search for elements
node := list.Search(2)
fmt.Println(node.Data)  // Output: 2

// Delete an element
list.DeleteNode(2)
list.Display()  // Output: 1 -> 3

// Get length
fmt.Println(list.Length())  // Output: 2