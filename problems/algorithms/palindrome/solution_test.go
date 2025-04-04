/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"Madam", true},
		{"", true},  // Empty string is considered palindrome
		{"a", true}, // Single character is always palindrome
		{"ab", false},
		{"Able was I ere I saw Elba", true},
	}

	for _, tc := range testCases {
		result := IsPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
