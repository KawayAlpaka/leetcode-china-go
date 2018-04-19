package productexceptself

import (
	"testing"

	"../../../mcommon"
)

type Case struct {
	input  []int
	expect []int
}

var cases = []Case{
	Case{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	Case{[]int{1, 0}, []int{0, 1}},
	Case{[]int{0, 0}, []int{0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for i, c := range cases {
		result := productExceptSelf2(c.input)
		if !mcommon.IsEqual(result, c.expect) {
			t.Errorf("Error productExceptSelf ：index=%d,result=%v,expect=%v", i, result, c.expect)
		}
	}
}

func BenchmarkProductExceptSelf_B(b *testing.B) {
	for _, c := range cases {
		productExceptSelf(c.input)
	}
}

func BenchmarkProductExceptSelf_B2(b *testing.B) {
	for _, c := range cases {
		productExceptSelf2(c.input)
	}
}
