/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package palindrome

import (
	"strings"
	"unicode"
)

// IsPalindrome checks if a string is a palindrome
// It ignores case, spaces, and punctuation
func IsPalindrome(s string) bool {
	// Convert to lowercase and remove non-alphanumeric characters
	var cleaned strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(unicode.ToLower(char))
		}
	}

	cleanedStr := cleaned.String()
	length := len(cleanedStr)

	// Check if the string reads the same forward and backward
	for i := 0; i < length/2; i++ {
		if cleanedStr[i] != cleanedStr[length-1-i] {
			return false
		}
	}

	return true
}
