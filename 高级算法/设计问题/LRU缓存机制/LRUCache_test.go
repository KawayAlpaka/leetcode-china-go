package lrucache

import (
	"fmt"
	"testing"
)

type Case struct {
	capacity int
	steps    []Step
}
type Step struct {
	isGet  bool
	key    int
	value  int
	expect int
}

var cases = []Case{
	Case{
		2,
		[]Step{
			Step{false, 1, 1, 0},
			Step{false, 2, 2, 0},
			Step{true, 1, 0, 1},
			Step{false, 3, 3, 0},
			Step{true, 2, 0, -1},
			Step{false, 4, 4, 0},
			Step{true, 1, 0, -1},
			Step{true, 3, 0, 3},
			Step{true, 4, 0, 4},
		},
	},
	Case{
		2,
		[]Step{
			Step{false, 2, 1, 0},
			Step{false, 1, 1, 0},
			Step{true, 2, 0, 1},
			Step{false, 4, 1, 0},
			Step{true, 1, 0, -1},
			Step{true, 2, 0, 1},
		},
	},
	Case{
		2,
		[]Step{
			Step{false, 2, 1, 0},
			Step{false, 1, 1, 0},
			Step{false, 2, 3, 0},
			Step{false, 4, 1, 0},
			Step{true, 1, 0, -1},
			Step{true, 2, 0, 3},
		},
	},
	Case{
		2,
		[]Step{
			Step{true, 2, 0, -1},
			Step{false, 2, 6, 0},
			Step{true, 1, 0, -1},
			Step{false, 1, 5, 0},
			Step{false, 1, 2, 0},
			Step{true, 1, 0, 2},
			Step{true, 2, 0, 6},
		},
	},
}

func TestLRUCache(t *testing.T) {

	for i, c := range cases {
		fmt.Printf("case:%d \n", i)
		cache := Constructor(c.capacity)
		for j, s := range c.steps {
			if s.isGet {
				result := cache.Get(s.key)
				if result != s.expect {
					t.Errorf("Error LRUCache case=%d,step=%d,result=%v,expect=%v", i, j, result, s.expect)
				}
			} else {
				cache.Put(s.key, s.value)
			}
		}

	}

}
