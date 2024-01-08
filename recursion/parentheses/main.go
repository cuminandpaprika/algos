package parentheses

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// n=1
// "()"
// n=2
// ["()()", "(())"]

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack("", 0, 0, &result, n)
	return result
}

func backtrack(s string, left, right int, result *[]string, n int) {
	if len(s) == 2*n {
		*result = append(*result, s)
	}
	if left < n {
		s = s + "("
		backtrack(s, left+1, right, result, n)
		s = trim(s)
	}
	if right < left {
		s = s + ")"
		backtrack(s, left, right+1, result, n)
		// s = trim(s)
	}
}

func trim(s string) string {
	if len(s) > 1 {
		return s[:len(s)-1]
	}
	return ""
}
