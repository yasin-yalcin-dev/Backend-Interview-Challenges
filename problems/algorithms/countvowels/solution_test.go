/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/algorithms/countvowels/solution_test.go
package countvowels

import "testing"

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"hello", 2},
		{"Beautiful", 5},
		{"Sky", 0},
		{"AEIOU", 5},
		{"aeiou", 5},
		{"", 0},
		{"bcdfghjklmnpqrstvwxyz", 0},
		{"Rhythm", 0},
		{"Quick brown fox", 3},
	}

	for _, tc := range testCases {
		result := CountVowels(tc.input)
		if result != tc.expected {
			t.Errorf("CountVowels(%q) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}
