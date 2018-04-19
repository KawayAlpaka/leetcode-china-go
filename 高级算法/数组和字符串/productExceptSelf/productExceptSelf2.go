//https://blog.csdn.net/wzy_1988/article/details/46916179
package productexceptself

func productExceptSelf2(nums []int) []int {
	multiply(nums, 1, 0, len(nums))

	return nums
}

func multiply(a []int, fwdProduct int, indx int, N int) int {
	revProduct := 1
	if indx < N {
		revProduct = multiply(a, fwdProduct*a[indx], indx+1, N)
		cur := a[indx]
		a[indx] = fwdProduct * revProduct
		revProduct *= cur
	}
	return revProduct
}
