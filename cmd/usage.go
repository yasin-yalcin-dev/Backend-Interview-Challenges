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
	fmt.Println("  list             - List all available problems")
	fmt.Println("\nExamples:")
	fmt.Println("  interview-challenges algorithms stringreversal \"hello world\"")
	fmt.Println("  interview-challenges algorithms twosum \"[2,7,11,15]\" 9")
	fmt.Println("  interview-challenges list")
}
