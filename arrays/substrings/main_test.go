package substrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthLongestSubstringWithoutRepeating(t *testing.T) {
	expected := []int{
		3,
		3,
	}
	input := []string{
		"abc",
		"abcaabcd",
	}

	for i, in := range input {
		t.Run("", func(t *testing.T) {
			result := LengthLongestSubstringWithoutRepeating(in)
			assert.Equal(t, result, expected[i])
		})
	}

}


