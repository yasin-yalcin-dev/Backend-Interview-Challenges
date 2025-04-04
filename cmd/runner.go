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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/countvowels"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/firstrepeatingcharacter"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/fizzbuzz"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/palindrome"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/stringreversal"
	"github.com/yasin-yalcin-dev/Backend-Interview-Challenges/problems/algorithms/twosum"
)

func RunAlgorithmProblem(problem string, args []string) {
	switch problem {
	case "stringreversal":
		runStringReversal(args)
	case "palindrome":
		runPalindrome(args)
	case "countvowels":
		runCountVowels(args)
	case "fizzbuzz":
		runFizzBuzz(args)
	case "twosum":
		runTwoSum(args)
	case "firstrepeatingcharacter":
		runFirstRepeatingCharacter(args)
	default:
		fmt.Printf("Unknown problem: %s\n", problem)
		fmt.Println("Use 'interview-challenges list' to see available problems")
	}
}

func runStringReversal(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: interview-challenges algorithms stringreversal <string>")
		return
	}
	input := args[0]
	result := stringreversal.ReverseString(input)
	fmt.Printf("Original: %s\nReversed: %s\n", input, result)
}

func runPalindrome(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: interview-challenges algorithms palindrome <string>")
		return
	}
	input := args[0]
	result := palindrome.IsPalindrome(input)
	fmt.Printf("String: %s\nIs Palindrome: %v\n", input, result)
}

func runCountVowels(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: interview-challenges algorithms countvowels <string>")
		return
	}
	input := args[0]
	result := countvowels.CountVowels(input)
	fmt.Printf("String: %s\nVowel Count: %d\n", input, result)
}

func runFizzBuzz(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: interview-challenges algorithms fizzbuzz <number>")
		return
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid number: %s\n", args[0])
		return
	}
	result := fizzbuzz.FizzBuzz(n)
	fmt.Printf("FizzBuzz up to %d:\n%s\n", n, strings.Join(result, ", "))
}

func runTwoSum(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: interview-challenges algorithms twosum \"[num1,num2,...]\" <target>")
		return
	}

	// Parse array from string
	arrayStr := args[0]
	// Try to parse as JSON first
	var nums []int
	err := json.Unmarshal([]byte(arrayStr), &nums)

	// If JSON parsing fails, try manual parsing
	if err != nil {
		arrayStr = strings.Trim(arrayStr, "[]")
		numStrs := strings.Split(arrayStr, ",")

		nums = make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				fmt.Printf("Invalid number in array: %s\n", numStr)
				return
			}
			nums[i] = num
		}
	}

	// Parse target
	target, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Invalid target: %s\n", args[1])
		return
	}

	result := twosum.TwoSum(nums, target)
	fmt.Printf("Array: %v\nTarget: %d\nIndices: %v\n", nums, target, result)
}

func runFirstRepeatingCharacter(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: interview-challenges algorithms firstrepeatingcharacter <string>")
		return
	}
	input := args[0]
	result := firstrepeatingcharacter.FirstRepeatingCharacter(input)
	if result == "" {
		fmt.Printf("String: %s\nNo repeating character found\n", input)
	} else {
		fmt.Printf("String: %s\nFirst Repeating Character: %s\n", input, result)
	}
}
