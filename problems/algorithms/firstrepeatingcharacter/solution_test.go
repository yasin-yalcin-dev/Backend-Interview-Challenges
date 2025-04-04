/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/algorithms/firstrepeatingcharacter/solution_test.go
package firstrepeatingcharacter

import "testing"

func TestFirstRepeatingCharacter(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"abcdefghija", "a"},
		{"hello", "l"},
		{"abcdef", ""},
		{"", ""},
		{"aabbcc", "a"},
		{"ABCBA", "B"},
		{"abba", "b"},
		{"abc123ab", "a"},
		{"zzz", "z"},
	}

	for _, tc := range testCases {
		result := FirstRepeatingCharacter(tc.input)
		if result != tc.expected {
			t.Errorf("FirstRepeatingCharacter(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
