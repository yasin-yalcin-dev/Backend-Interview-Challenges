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

func ShowUsage() {
	fmt.Println("\nUsage: interview-challenges <category> <problem> [arguments...]")
	fmt.Println("\nCategories:")
	fmt.Println("  algorithms       - Run algorithm problems")
	fmt.Println("  oop              - Run object-oriented design problems")
	fmt.Println("  datastructures   - Run data structure implementations")
	fmt.Println("  list             - List all available problems")
	fmt.Println("\nExamples:")
	fmt.Println("  interview-challenges algorithms stringreversal \"hello world\"")
	fmt.Println("  interview-challenges algorithms twosum \"[2,7,11,15]\" 9")
	fmt.Println("  interview-challenges oop shapehierarchy")
	fmt.Println("  interview-challenges datastructures linkedlist")
	fmt.Println("  interview-challenges list")
}
