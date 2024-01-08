package substrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthLongestSubstringWithoutRepeating(t *testing.T) {
	expected := []int{
		0,
		3,
		4,
		4,
		4,
		3,
	}
	input := []struct {
		Name  string
		Input string
	}{
		{Name: "Empty", Input: ""},
		{Name: "Unique", Input: "abc"},
		{Name: "RepeatMiddle", Input: "abcaabcd"},
		{Name: "RepeatEnd", Input: "abcdaaa"},
		{Name: "RepeatStart", Input: "aaabcd"},
		{Name: "Testcase", Input: "dbdf"},
	}

	for i, in := range input {
		t.Run(in.Name, func(t *testing.T) {
			result := LengthLongestSubstringWithoutRepeating(in.Input)
			assert.Equal(t, expected[i], result)
		})
	}

}
