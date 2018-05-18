package threesum

import "../../../mcommon"

func threeSum(nums []int) [][]int {
	len := len(nums)
	var r [][]int
	for x := 0; x < len-2; x++ {
		for y := x + 1; y < len-1; y++ {
			xy := nums[x] + nums[y]
			for z := y + 1; z < len; z++ {
				if xy+nums[z] == 0 {
					r = append(r, []int{nums[x], nums[y], nums[z]})
				}
			}
		}
	}
	r = mcommon.Distinct2WArr(r)
	return r
}
