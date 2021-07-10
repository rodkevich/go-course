

// CacheMap ...
type CacheMap map[int64]int64

// NewCacheMap ...
func NewCacheMap() CacheMap {
	c := CacheMap{0: 0, 1: 1}
	return c
}

// SizedSequence64bitUseCacheHardWay returns a new slice of fibonacci numbers with required length if possible. Breaks on 64 bit overflow
func SizedSequence64bitUseCacheHardWay(size int64, cache CacheMap) ([]int64, error) {
	s64 := SizedSequence64bit
	if _, ok := cache[size]; ok {
		cachedArray := make([]int64, 0)
		//cachedArray := make([]int64, len(cache))

		for i := cache[0]; i <= cache[size]; i++ {
			for _, v := range cache {
				cachedArray = append(cachedArray, v)
			}
		}
		fmt.Println("\nUsing cache")
		fmt.Println(cachedArray)
		fmt.Println("CachedArray")
		//fmt.Printf("\n%v", cachedArray)
		return cachedArray, nil
	} else {
		arr, err := s64(size)
		if err != nil {
			return nil, err
		}
		return arr, err
	}

}
