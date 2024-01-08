package substrings

// Given a string s, find the length of the longest substring without repeating characters.

func LengthLongestSubstringWithoutRepeating(s string) int {
	longestWindow := 0
	startIndex := 0
	seenChars := map[rune]int{}

	// dvdf
	for i, char := range s {
		if position := seenChars[char]; position > startIndex {
			// Shift the left side of the window
			startIndex = position
		}

		// Keep track of where we need to seek to, to remove the duplicate
		seenChars[char] = i + 1

		if currentWindow := i - startIndex + 1; currentWindow > longestWindow {
			longestWindow = currentWindow
		}

	}

	return longestWindow
}
