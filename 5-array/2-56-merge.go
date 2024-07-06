package array

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	left := intervals[0][0]
	right := intervals[0][1]
	for index := range intervals {
		if right >= intervals[index][0] {
			right = max(right, intervals[index][1])
		} else {
			res = append(res, []int{left, right})
			left = intervals[index][0]
			right = intervals[index][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
