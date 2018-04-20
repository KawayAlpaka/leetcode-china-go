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
	// fmt.Printf("start get:key:%d \n", key)
	if v, ok := cache.m[key]; ok {
		cache.moveToTop(key)
		return v
	}
	// fmt.Printf("end get:key:%d \n", key)
	// fmt.Printf("m:%v \n", cache.m)
	// fmt.Printf("cache:%v \n", cache.cache)

	return -1
}

// Put ...
func (cache *LRUCache) Put(key int, value int) {
	// fmt.Printf("start put:key:%d,value:%d \n", key, value)
	if _, ok := cache.m[key]; ok {
		cache.moveToTop(key)
	} else {
		cache.cache = append(cache.cache, key)
		outLen := len(cache.cache) - cache.capacity
		if outLen > 0 {
			// ks := cache.cache[:outLen]
			// cache.cache = cache.cache[outLen:]
			// for _, k := range ks {
			// 	delete(cache.m, k)
			// }
			k := cache.cache[0]
			cache.cache = cache.cache[1:]
			delete(cache.m, k)
			// if k != key {
			// 	delete(cache.m, k)
			// }
		}
	}
	cache.m[key] = value

	// fmt.Printf("end put:key:%d,value:%d \n", key, value)
	// fmt.Printf("m:%v \n", cache.m)
	// fmt.Printf("cache:%v \n", cache.cache)

}
func (cache *LRUCache) moveToTop(key int) {
	l1 := len(cache.cache)
	for i := 0; 1 < l1; i++ {
		if cache.cache[i] == key {
			cache.cache = append(cache.cache[:i], cache.cache[i+1:]...)
			break
		}
	}
	cache.cache = append(cache.cache, key)
}
