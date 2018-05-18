package mcommon

import (
	"sort"
	"strconv"
	"strings"
)

// IsEqual 对比两个整形数组是否相同
func IsEqual(nums1 []int, nums2 []int) bool {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

// Is2Equal 对比两个二维整形数组是否相同
func Is2Equal(nums1 [][]int, nums2 [][]int) bool {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		l11 := len(nums1[i])
		l21 := len(nums2[i])
		if l11 != l21 {
			return false
		}
		for j := 0; j < l2; j++ {
			if nums1[i][j] != nums2[i][j] {
				return false
			}
		}
	}
	return true
}

// Create2WIntArray 创建一个空的整形二维数组
func Create2WIntArray(rl int, cl int) [][]int {
	var r [][]int
	for i := 0; i < rl; i++ {
		r = append(r, make([]int, cl))
	}
	return r
}

// Distinct2WArr 二维数组去重 规则:不考虑顺序，即 [-1,1] 与 [1,-1]重复
func Distinct2WArr(arr [][]int) [][]int {
	// rowLen := len(arr)
	m := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		var intSlice sort.IntSlice
		intSlice = arr[i]
		intSlice.Sort()
		strArr := IntArr2StrArr(intSlice)
		key := strings.Join(strArr, "a")
		m[key]++
		if m[key] > 1 {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
	}
	// fmt.Printf("%v \n", m)
	// fmt.Printf("%v \n", arr)
	return arr
}

// IntArr2StrArr 数字数组换成
func IntArr2StrArr(arr []int) []string {
	r := []string{}
	for i := 0; i < len(arr); i++ {
		r = append(r, strconv.Itoa(arr[i]))
	}
	return r
}
