/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package main

import (
	"fmt"

	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/datastructures/binarysearchtree"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/datastructures/graph"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/datastructures/hashtable"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/datastructures/linkedlist"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/datastructures/stackqueue"
)

// RunDataStructureProblem runs the specified data structure problem
func RunDataStructureProblem(problem string, args []string) {
	switch problem {
	case "linkedlist":
		runLinkedList()
	case "stackqueue":
		runStackQueue()
	case "binarysearchtree":
		runBinarySearchTree()
	case "hashtable":
		runHashTable()
	case "graph":
		runGraph()
	default:
		fmt.Printf("Unknown problem: %s\n", problem)
		fmt.Println("Use 'interview-challenges list' to see available problems")
	}
}

func runLinkedList() {
	fmt.Println("Running Linked List example...")
	linkedlist.RunExample()
}

func runStackQueue() {
	fmt.Println("Running Stack & Queue example...")
	stackqueue.RunExample()
}

func runBinarySearchTree() {
	fmt.Println("Running Binary Search Tree example...")
	binarysearchtree.RunExample()
}

func runHashTable() {
	fmt.Println("Running Hash Table example...")
	hashtable.RunExample()
}

func runGraph() {
	fmt.Println("Running Graph example...")
	graph.RunExample()
}
