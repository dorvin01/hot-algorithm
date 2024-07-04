package subsequence

func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	hmap := make(map[int]int)
	hmap[0] = 1
	for index := range nums {
		pre += nums[index]
		count += hmap[pre-k]
		hmap[pre] = max(hmap[pre], 0) + 1
	}
	return count
}
