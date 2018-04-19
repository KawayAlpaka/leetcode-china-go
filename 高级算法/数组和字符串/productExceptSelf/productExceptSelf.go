package productexceptself

func productExceptSelf(nums []int) []int {
	l1 := len(nums)
	output := make([]int, l1)
	output[l1-1] = 1
	for i := 1; i < l1; i++ {
		output[l1-1-i] = output[l1-i] * nums[l1-1-i]
	}
	nums[0] = 1
	for i := 1; i < l1; i++ {
		nums[i] = nums[i-1] * nums[i]
	}
	for i := 0; i < l1; i++ {
		output[i] = output[l1-1-i] * nums[l1-1-i]
	}
	return output
}
