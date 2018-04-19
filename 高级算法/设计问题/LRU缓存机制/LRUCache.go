package lrucache

// LRUCache ...
type LRUCache struct {
	capacity int
	cache    []int
	m        map[int]int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, []int{}, make(map[int]int)}
}

// Get ...
func (cache *LRUCache) Get(key int) int {
	if v, ok := cache.m[key]; ok {
		l1 := len(cache.cache)
		for i := 0; 1 < l1; i++ {
			if cache.cache[i] == key {
				cache.cache = append(cache.cache[:i], cache.cache[i+1:]...)
				break
			}
		}
		cache.cache = append(cache.cache, key)
		return v
	}
	return -1
}

// Put ...
func (cache *LRUCache) Put(key int, value int) {
	cache.m[key] = value
	cache.cache = append(cache.cache, key)
	outLen := len(cache.cache) - cache.capacity
	if outLen > 0 {
		ks := cache.cache[:outLen]
		cache.cache = cache.cache[outLen:]
		for _, k := range ks {
			delete(cache.m, k)
		}
	}
}
