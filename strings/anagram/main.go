package anagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func isAnagram(s string, t string) bool {
	charCount := make(map[rune]int)

	// Count chars in first string
	for _, char := range s {
		if _, ok := charCount[char]; ok {
			charCount[char]++
		} else {
			charCount[char] = 1
		}
	}

	for _, char := range t {
		if count, ok := charCount[char]; ok {
			// Remove the char if we only have one
			if count == 1 {
				delete(charCount, char)
			} else {
				charCount[char]--
			}
		} else {
			// If we can't find the char, false
			return false
		}
	}

	// Early return if we have unused chars
	if len(charCount) > 0 {
		return false
	}

	return true
}

// Edge cases
// - Empty string input
// - Subset string.

// Alternatives:
// - Sort strings and compare
