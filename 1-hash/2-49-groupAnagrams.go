package hash

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hmap := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		sortedStr := string(bytes)
		hmap[sortedStr] = append(hmap[sortedStr], str)
	}
	var result [][]string
	for _, value := range hmap {
		result = append(result, value)
	}
	return result
}
