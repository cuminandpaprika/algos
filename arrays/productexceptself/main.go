package productexceptself

// [1,2,3,4]
// [nums[1]*nums[2]*nums[3],nums[0]*nums[2]*nums[3], nums[0]*nums[1]*nums[3], nums[0]*nums[1]*nums[2]]
// [24, 12, 8, 6]
// [nums[0], nums[0]*nums[1], nums[0]*nums[1]*nums[2], nums[0]*nums[1]*nums[2]*nums[3]]
// [1, 2, 6, 24]

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	productSoFarLeft := make([]int, len(nums))
	productSoFarRight := make([]int, len(nums))

	for i, num := range nums {
		if i != 0 {
			productSoFarLeft[i] = productSoFarLeft[i-1] * num
		} else {
			productSoFarLeft[i] = nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 {
			productSoFarRight[i] = productSoFarRight[i+1] * nums[i]
		} else {
			productSoFarRight[i] = nums[i]
		}
	}

	for i := range nums {
		if i == 0 {
			result[i] = productSoFarRight[i+1]
		} else if i == len(nums)-1 {
			result[i] = productSoFarLeft[i-1]
		} else {
			result[i] = productSoFarLeft[i-1] * productSoFarRight[i+1]
		}
	}

	return result
}
