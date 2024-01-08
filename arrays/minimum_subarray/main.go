package minimum_subarray

// Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.
// Example 2:

// Input: target = 4, nums = [1,4,4]
// Output: 1
// Example 3:

// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// Constraints:

// 1 <= target <= 10^9
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^4

// [1,2,3,4]
// [1,3,6,10]
// From 1 to 3 is 10-1
// From 2 to 3 is 10-3

// Questions
// Is the input always positive?
// Is the input ever overflow?
//

func MinimumSubArray(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	// Impossible to reach number
	smallestWindowSize := len(nums) + 2

	sum := make([]int, len(nums))

	for i, input := range nums {
		if i == 0 {
			sum[i] = input
		} else {
			// Subtract left items
			sum[i] = sum[i-1] + input
		}

		for left := 0; left <= i; left++ {
			currentSum := 0
			currentSumSize := 0
			// Window of size 1
			if left == i {
				currentSum = input
				currentSumSize = 1
			} else if left == 0 {
				currentSum = sum[i]
				currentSumSize = i + 1
			} else {
				// All other cases
				currentSum = sum[i] - sum[left]
				currentSumSize = i - left
			}

			if currentSum >= target {
				// Record only if larger
				if smallestWindowSize > currentSumSize {
					smallestWindowSize = currentSumSize
				}
			} else {
				// No point making window smaller
				break
			}

		}
	}

	if smallestWindowSize == len(nums)+2 {
		return 0
	}

	return smallestWindowSize
}
