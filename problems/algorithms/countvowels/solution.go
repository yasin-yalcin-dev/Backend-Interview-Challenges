/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package countvowels

import "strings"

// CountVowels counts the number of vowels in a given string
func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
