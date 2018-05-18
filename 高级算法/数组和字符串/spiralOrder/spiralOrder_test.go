package spiralorder

import (
	"testing"

	"../../../mcommon"
)

type Case struct {
	input  [][]int
	expect []int
}

var cases = []Case{
	Case{[][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	Case{[][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	Case{[][]int{},
		[]int{}},
	Case{[][]int{
		[]int{},
	},
		[]int{}},
}

func TestSpiralOrder(t *testing.T) {
	for i, c := range cases {
		result := spiralOrder(c.input)
		if !mcommon.IsEqual(result, c.expect) {
			t.Errorf("Error spiralOrder ：index=%d,result=%v,expect=%v \n", i, result, c.expect)
		}
	}
}
