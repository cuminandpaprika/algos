package subsets

func subsets(nums []int) [][]int {
	result := [][]int{}

	for i := 0; i < len(nums)+1; i++ {
		backtrack(0, i, []int{}, nums, &result)
	}

	return result
}

func backtrack(startIndex int, length int, curr []int, nums []int, result *[][]int) {
	if len(curr) == length {
		*result = append(*result, append([]int{}, curr...))
		return
	}
	for i := startIndex; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack(i+1, length, curr, nums, result)
		curr = pop(curr)
	}

}

func pop(n []int) []int {
	if len(n) < 1 {
		return []int{}
	}
	return n[:len(n)-1]
}
