package longestconsecutive

func longestConsecutive(nums []int) int {
	l1 := len(nums)
	m := make(map[int]int)
	r := make(map[int]int)
	// max := 0
	// min := 0
	for i := 0; i < l1; i++ {
		m[nums[i]] = 1
		// if nums[i] >= max {
		// 	max = nums[i]
		// }
		// if nums[i] <= min {
		// 	min = nums[i]
		// }
	}
	// fmt.Printf("m:%v \n", m)
	// fmt.Printf("min:%d,max:%d \n", min, max)
	// for i := min; i <= max; i++ {
	// 	_, ok1 := m[i]
	// 	if ok1 {
	// 		// fmt.Printf("num1:%d \n", i)
	// 		num2, ok2 := r[i-1]
	// 		if ok2 {
	// 			r[i] = num2 + 1
	// 		} else {
	// 			r[i] = 1
	// 		}
	// 	}
	// }
	for k := range m {
		num3, ok3 := r[k-1]
		if ok3 {
			r[k] = num3 + 1
		} else {
			r[k] = 1
		}
		for i := 1; true; i++ {
			if num4, ok4 := r[k+i]; ok4 {
				if num4 != num3+i+1 {
					r[k+i] = num3 + i + 1
				}
			} else {
				break
			}
		}

	}

	// fmt.Printf("r:%v \n", r)
	long := 0
	for _, v := range r {
		if v > long {
			long = v
		}
	}
	return long
}
