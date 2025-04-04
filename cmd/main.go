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
	"os"
)

func main() {
	fmt.Println("Backend Interview Challenges Runner")

	if len(os.Args) < 2 {
		ShowUsage()
		return
	}

	category := os.Args[1]

	if category == "list" {
		ListProblems()
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Please specify a problem to run")
		ShowUsage()
		return
	}

	problem := os.Args[2]
	args := os.Args[3:]

	switch category {
	case "algorithms":
		RunAlgorithmProblem(problem, args)
	case "oop":
		fmt.Println("OOP problems are not implemented yet")
	case "datastructures":
		fmt.Println("Data structure problems are not implemented yet")
	case "systemdesign":
		fmt.Println("System design problems are not implemented yet")
	default:
		fmt.Printf("Unknown category: %s\n", category)
		ShowUsage()
	}
}
