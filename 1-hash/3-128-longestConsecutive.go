package hash

func longestConsecutive(nums []int) int {
	hmap := make(map[int]struct{})
	for _, num := range nums {
		hmap[num] = struct{}{}
	}
	count := 0
	res := 0
	for num := range hmap {
		if _, ok := hmap[num-1]; !ok {
			// 头节点
			count = 1
			curr := num + 1
			for {
				if _, ok := hmap[curr]; !ok {
					break
				}
				count++
				curr++
			}
			// 不满足条件时候更新res
			res = max(res, count)
		}
	}
	return res
}
