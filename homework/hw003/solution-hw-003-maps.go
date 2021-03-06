package hw003

import (
	"fmt"
	"sort"
)

// PrintValuesSortedByIncrKeys prints map values sorted in order of increasing keys
func PrintValuesSortedByIncrKeys(m map[int]string) {
	// allocate slices with length according to map
	rtn := make([]string, 0, len(m)) // to return results
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	// sort map keys alphabetically
	sort.Ints(keys)

	for _, key := range keys {
		rtn = append(rtn, m[key])
	}

	fmt.Printf("\nresult | PrintValuesSortedByIncrKeys() val=%v", rtn)
	return
}
