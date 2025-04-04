/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package main

import "fmt"

func ListProblems() {
	fmt.Println("\nAvailable problems:")

	fmt.Println("\nALGORITHMS:")
	fmt.Println("  1. stringreversal        - Reverse a string")
	fmt.Println("  2. palindrome            - Check if a string is a palindrome")
	fmt.Println("  3. countvowels           - Count vowels in a string")
	fmt.Println("  4. fizzbuzz              - FizzBuzz implementation")
	fmt.Println("  5. twosum                - Find indices of two numbers that add up to target")
	fmt.Println("  6. firstrepeatingcharacter - Find first repeating character in a string")

	fmt.Println("\nOOP DESIGN:")
	fmt.Println("  1. shapehierarchy        - Implement a polymorphic shape class hierarchy")
	fmt.Println("  2. bankingsystem         - Design classes for a banking application")
	fmt.Println("  3. factorypattern        - Implement payment methods using the Factory pattern")
	fmt.Println("  4. observerpattern       - Create a newsletter subscription system")
	fmt.Println("  5. singleton             - Database connection pool with dependency injection")

	fmt.Println("\nDATA STRUCTURES:")
	fmt.Println("  1. linkedlist            - Implement a singly linked list")
	fmt.Println("  2. stackqueue            - Implement stack and queue data structures")
	fmt.Println("  3. binarysearchtree      - Implement a binary search tree")
	fmt.Println("  4. hashtable             - Implement a hash table")
	fmt.Println("  5. graph                 - Implement a graph with common algorithms")

	fmt.Println("\nMore categories coming soon...")
}
