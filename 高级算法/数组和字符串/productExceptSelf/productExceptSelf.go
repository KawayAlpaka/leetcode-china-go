package productexceptself

func productExceptSelf(nums []int) []int {
	l1 := len(nums)
	output := make([]int, l1)
	output[l1-1] = 1
	for i := 1; i < l1; i++ {
		output[l1-1-i] = output[l1-i] * nums[l1-i]
	}
	s := 1
	for i := 0; i < l1; i++ {
		_n := nums[i]
		nums[i] = s
		s = s * _n
	}
	for i := 0; i < l1; i++ {
		output[i] = output[i] * nums[i]
	}
	return output
}
