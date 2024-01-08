package combinations

import "fmt"

// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
// n=2, k=1
// [[1],[2]]
// n=3, k=2
// [[1,2],[1,3]

func combine(n int, k int) [][]int {
	result := [][]int{}
	backtrack([]int{}, 1, n, k, &result)
	return result
}

func backtrack(curr []int, startIndex int, n int, k int, result *[][]int) {
	fmt.Printf("Curr: %d, startIndex: %d , result: %d \n", curr, startIndex, result)
	// We've got enough
	if len(curr) == k {
		*result = append(*result, append([]int{}, curr...))
		return
	}
	for i := startIndex; i <= n; i++ {
		curr = append(curr, i)
		backtrack(curr, i+1, n, k, result)
		curr = trim(curr)
	}
}

func trim(n []int) []int {
	if len(n) == 1 {
		return []int{}
	}
	return n[:len(n)-1]
}
