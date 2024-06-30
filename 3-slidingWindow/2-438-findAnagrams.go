package slidingWindow

import "reflect"

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	target := getCharCount(p, 0, len(p))
	curr := getCharCount(s, 0, len(p)-1)
	for i := 0; i < len(s)-len(p)+1; i++ {
		curr[int(s[i+len(p)-1])-int('a')]++
		if reflect.DeepEqual(target, curr) {
			res = append(res, i)
		}
		curr[int(s[i])-int('a')]--
	}
	return res
}

func getCharCount(target string, start, end int) [26]int {
	var res [26]int
	for i := start; i < end; i++ {
		position := int(target[i]) - int('a')
		res[position]++
	}
	return res
}
