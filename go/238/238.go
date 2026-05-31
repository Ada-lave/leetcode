package task_238



func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))

	left := make([]int, len(nums))
	right := make([]int, len(nums))
	
	prefix := 1
	suffix := 1

	for i := 0; i < len(nums); i++ {
		left[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = suffix
		suffix *= nums[i]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}

	return result
}