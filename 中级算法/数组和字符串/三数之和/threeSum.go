package threesum

import (
	"fmt"
	"sort"

	"../../../mcommon"
)

func threeSum(nums []int) [][]int {
	r := [][]int{}
	l := len(nums)
	fmt.Printf("len: %d \n", l)
	if l < 3 {
		return [][]int{}
	}
	var intSlice sort.IntSlice = nums
	intSlice.Sort()
	intSlice = mcommon.RemoveDuplicates(intSlice, 3)
	l = len(intSlice)
	fmt.Printf("RemoveDuplicates len: %d \n", l)
	if l < 3 {
		return [][]int{}
	}
	// if l < 3 {
	// 	intSlice = append(intSlice, intSlice[l-1], intSlice[l-1])
	// 	l = len(intSlice)
	// }
	fmt.Printf("%v \n", intSlice)

	for i := 1; i < l-1; i++ {
		low := 0
		high := l - 1
		target := -intSlice[i]

		for true {
			// fmt.Printf("low=%d,i=%d,higt=%d \n", low, i, high)
			if low >= i || high <= i {
				break
			}
			var s = intSlice[low] + intSlice[high]
			if s == target {
				r = append(r, []int{intSlice[low], intSlice[i], intSlice[high]})
				high--
				low++
			}
			if s > target {
				high--
				continue
			}
			if s < target {
				low++
				continue
			}

		}
	}
	r = mcommon.Distinct2WArr(r)
	return r
}

func threeSum2(nums []int) [][]int {
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
