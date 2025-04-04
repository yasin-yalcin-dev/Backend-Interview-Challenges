/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package twosum

// TwoSum returns the indices of two numbers such that they add up to target
func TwoSum(nums []int, target int) []int {
	// Create a map to store values and their indices
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement (the value we need to reach target)
		complement := target - num

		// Check if the complement exists in our map
		if j, found := numMap[complement]; found {
			return []int{j, i} // Return the indices
		}

		// Store current number and its index
		numMap[num] = i
	}

	// No solution found (though the problem states there will always be one)
	return []int{}
}
