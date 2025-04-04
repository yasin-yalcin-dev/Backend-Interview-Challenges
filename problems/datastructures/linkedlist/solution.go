/*
** ** ** ** ** **

	\ \ / / \ \ / /
	 \ V /   \ V /
	  | |     | |
	  |_|     |_|
	 Yasin   Yalcin
*/
package linkedlist

import "fmt"

// RunExample demonstrates the linked list implementation
func RunExample() {
	fmt.Println("Linked List Example:")
	fmt.Println("-------------------")

	// Create a new linked list
	list := New()

	// Check if it's empty
	fmt.Printf("Is the list empty? %v\n", list.IsEmpty())

	// Insert elements
	fmt.Println("\nInserting elements:")
	fmt.Println("1. Insert 10 at the beginning")
	list.InsertAtBeginning(10)
	list.Display()

	fmt.Println("2. Insert 20 at the end")
	list.InsertAtEnd(20)
	list.Display()

	fmt.Println("3. Insert 30 at the end")
	list.InsertAtEnd(30)
	list.Display()

	fmt.Println("4. Insert 15 after 10")
	list.InsertAfter(10, 15)
	list.Display()

	fmt.Println("5. Insert 25 after 20")
	list.InsertAfter(20, 25)
	list.Display()

	// Get list information
	fmt.Printf("\nList length: %d\n", list.Length())

	// Search for elements
	fmt.Println("\nSearching for elements:")
	node := list.Search(15)
	if node != nil {
		fmt.Printf("Found node with value: %v\n", node.Data)
	} else {
		fmt.Println("Node not found")
	}

	node = list.Search(100)
	if node != nil {
		fmt.Printf("Found node with value: %v\n", node.Data)
	} else {
		fmt.Println("Node with value 100 not found")
	}

	// Delete elements
	fmt.Println("\nDeleting elements:")
	fmt.Println("1. Delete node with value 15")
	list.DeleteNode(15)
	list.Display()

	fmt.Println("2. Delete node with value 30")
	list.DeleteNode(30)
	list.Display()

	// Get head and tail
	head := list.GetHead()
	tail := list.GetTail()
	fmt.Printf("\nHead value: %v\n", head.Data)
	fmt.Printf("Tail value: %v\n", tail.Data)

	// Convert to slice
	fmt.Println("\nConverting list to slice:")
	slice := list.ToSlice()
	fmt.Printf("Slice: %v\n", slice)

	// Create from slice
	fmt.Println("\nCreating new list from slice [5, 6, 7, 8]:")
	newList := FromSlice([]interface{}{5, 6, 7, 8})
	newList.Display()

	// Reverse the list
	fmt.Println("\nReversing the list:")
	newList.Reverse()
	newList.Display()
}
