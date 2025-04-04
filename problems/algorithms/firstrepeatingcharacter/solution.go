/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package firstrepeatingcharacter

// FirstRepeatingCharacter finds the first character that repeats in a string
// Returns an empty string if no character repeats
func FirstRepeatingCharacter(s string) string {
	// Use a map to track seen characters
	charMap := make(map[rune]bool)

	// Iterate through the string
	for _, char := range s {
		// If character already seen, return it
		if charMap[char] {
			return string(char)
		}

		// Mark character as seen
		charMap[char] = true
	}

	// No repeating character found
	return ""
}
