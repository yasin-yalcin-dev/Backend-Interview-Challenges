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

	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/oop/bankingsystem"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/oop/factorypattern"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/oop/observerpattern"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/oop/shapehierarchy"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/oop/singleton"
)

// RunOOPProblem runs the specified OOP design problem
func RunOOPProblem(problem string, args []string) {
	switch problem {
	case "shapehierarchy":
		runShapeHierarchy()
	case "bankingsystem":
		runBankingSystem()
	case "factorypattern":
		runFactoryPattern()
	case "observerpattern":
		runObserverPattern()
	case "singleton":
		runSingleton()
	default:
		fmt.Printf("Unknown problem: %s\n", problem)
		fmt.Println("Use 'interview-challenges list' to see available problems")
	}
}

func runShapeHierarchy() {
	fmt.Println("Running Shape Hierarchy example...")
	shapehierarchy.RunExample()
}

func runBankingSystem() {
	fmt.Println("Running Banking System example...")
	bankingsystem.RunExample()
}

func runFactoryPattern() {
	fmt.Println("Running Factory Pattern example...")
	factorypattern.RunExample()
}

func runObserverPattern() {
	fmt.Println("Running Observer Pattern example...")
	observerpattern.RunExample()
}

func runSingleton() {
	fmt.Println("Running Singleton & Dependency Injection example...")
	singleton.RunExample()
}
