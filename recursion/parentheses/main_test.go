package parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "one", n: 1, want: []string{"()"}},
		{name: "two", n: 2, want: []string{"(())", "()()"}},
		{name: "three", n: 3, want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{name: "four", n: 4, want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
