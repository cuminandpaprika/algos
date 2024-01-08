package palindrome

import (
	"math"
	"regexp"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.
func isPalindrome(s string) bool {
	clean := cleanString(s)
	return isSymmetrical(clean)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9]+`)

// cleanString takes a string and removes all non alphanumerical characters.
// It also converts all letters to lowercase.
func cleanString(s string) string {
	lowerCase := strings.ToLower(s)
	return nonAlphanumericRegex.ReplaceAllString(lowerCase, "")
}

// isSymmetrical returns true if a string is symmetrical.
// Handles odd and even length strings
// Len: 2, pivot = 1
// Len: 3, pivot = 1
func isSymmetrical(s string) bool {
	pivot := int(math.Floor(float64(len(s) / 2)))
	for i := 0; i < pivot; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
