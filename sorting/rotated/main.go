package rotated

// We can't use binary search for an unordered slice
// [4,5,6,7,0,1,2]
func search(nums []int, target int) int {
	// find the pivot. AKA Find smallest element
	pivot := findPivot(nums)

	if nums[pivot] == target {
		return pivot
	}

	// No rotation. Full window
	if pivot == 0 {
		return findElement(0, len(nums)-1, nums, target)
	}

	if target < nums[0] {
		return findElement(pivot, len(nums)-1, nums, target)
	} else {
		return findElement(0, pivot, nums, target)
	}
}

func findPivot(nums []int) int {
	left := 0
	right := len(nums) - 1

	// Handle len 1
	if len(nums) == 1 {
		return 0
	}

	// Handle unrotated case
	if nums[left] < nums[right] {
		return 0
	}

	for left <= right {
		midPoint := (right + left) / 2
		if nums[midPoint] > nums[midPoint+1] {
			return midPoint + 1
		} else {
			if nums[midPoint] < nums[left] {
				right = midPoint - 1
			} else {
				left = midPoint + 1
			}
		}
	}
	return 0
}

func findElement(left, right int, nums []int, target int) int {
	// Handle length 1
	if len(nums) == 1 && nums[0] != target {
		return -1
	}

	for left <= right {
		midpoint := (right + left) / 2
		if nums[midpoint] == target {
			return midpoint
		}

		if nums[midpoint] > target {
			// Search to left
			right = midpoint - 1
		} else {
			// Search to right
			left = midpoint + 1
		}
	}
	return -1
}
