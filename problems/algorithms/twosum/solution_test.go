/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{1, 5, 8, 3, 9, 2}, 10, []int{0, 4}},
	}

	for _, tc := range testCases {
		result := TwoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TwoSum(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
		}
	}
}
