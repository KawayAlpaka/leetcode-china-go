package lrucache

import "testing"

type Case struct {
	isGet  bool
	key    int
	value  int
	expect int
}

var cases1 = []Case{
	Case{false, 1, 1, 0},
	Case{false, 2, 2, 0},
	Case{true, 1, 0, 1},
	Case{false, 3, 3, 0},
	Case{true, 2, 0, -1},
	Case{false, 4, 4, 0},
	Case{true, 1, 0, -1},
	Case{true, 3, 0, 3},
	Case{true, 4, 0, 4},
}
var cases2 = []Case{
	Case{false, 2, 1, 0},
	Case{false, 1, 1, 0},
	Case{true, 2, 0, 1},
	Case{false, 4, 1, 0},
	Case{true, 1, 0, -1},
	Case{true, 2, 0, 1},
}
var cases3 = []Case{
	Case{false, 2, 1, 0},
	Case{false, 1, 1, 0},
	Case{false, 2, 3, 0},
	Case{false, 4, 1, 0},
	Case{true, 1, 0, -1},
	Case{true, 2, 0, 3},
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	for i, c := range cases1 {
		if c.isGet {
			result := cache.Get(c.key)
			if result != c.expect {
				t.Errorf("Error LRUCache ：index=%d,result=%v,expect=%v", i, result, c.expect)
			}
		} else {
			cache.Put(c.key, c.value)
		}
	}
	cache = Constructor(2)
	for i, c := range cases2 {
		if c.isGet {
			result := cache.Get(c.key)
			if result != c.expect {
				t.Errorf("Error LRUCache ：index=%d,result=%v,expect=%v", i, result, c.expect)
			}
		} else {
			cache.Put(c.key, c.value)
		}
	}
	cache = Constructor(2)
	for i, c := range cases3 {
		if c.isGet {
			result := cache.Get(c.key)
			if result != c.expect {
				t.Errorf("Error LRUCache ：index=%d,result=%v,expect=%v", i, result, c.expect)
			}
		} else {
			cache.Put(c.key, c.value)
		}
	}
}
