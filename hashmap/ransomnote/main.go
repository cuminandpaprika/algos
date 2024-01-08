package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	letters := map[rune]int{}
	for _, letter := range magazine {
		if _, ok := letters[letter]; ok {
			letters[letter]++
		} else {
			letters[letter] = 1
		}

	}

	for _, letter := range ransomNote {
		if count, ok := letters[letter]; !ok || count < 1 {
			return false
		}
		letters[letter]--
	}

	return true
}
