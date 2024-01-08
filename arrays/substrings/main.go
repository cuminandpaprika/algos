package substrings

// Given a string s, find the length of the longest substring without repeating characters.

func LengthLongestSubstringWithoutRepeating(s string) int {
	longestWindow := 0
	startIndex := 0
	endIndex := 0
	seenChars := map[rune]struct{}{}

	if len(s) == 0 {
		return 0
	}
	// dvdf
	for i, char := range s {
		if _, ok := seenChars[char]; ok {
			// Reset window
			startIndex = i
			endIndex = i
			seenChars = map[rune]struct{}{}
		} else {
			// Increment window size
			endIndex = i

			// Update longest window
			if endIndex-startIndex > longestWindow {
				longestWindow = endIndex - startIndex
			}
		}
		// Add seen rune to map
		seenChars[char] = struct{}{}

	}

	return longestWindow + 1
}
