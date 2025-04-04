/*
** ** ** ** ** **

	\ \ / / \ \ / /
	 \ V /   \ V /
	  | |     | |
	  |_|     |_|
	 Yasin   Yalcin
*/
package stringreversal

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"golang", "gnalog"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"}, // palindrome
	}

	for _, tc := range testCases {
		result := ReverseString(tc.input)
		if result != tc.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
