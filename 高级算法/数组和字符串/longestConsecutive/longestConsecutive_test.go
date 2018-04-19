package longestconsecutive

import (
	"testing"
)

type Case struct {
	input  []int
	expect int
}

var cases = []Case{
	Case{[]int{100, 4, 200, 1, 3, 2},
		4},
	Case{[]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645},
		3},
}

func TestLongestConsecutive(t *testing.T) {
	for i, c := range cases {
		result := longestConsecutive(c.input)
		if result != c.expect {
			t.Errorf("Error longestConsecutive ：index=%d,result=%v,expect=%v", i, result, c.expect)
		}
	}
}
