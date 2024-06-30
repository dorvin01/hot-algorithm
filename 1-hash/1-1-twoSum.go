package hash

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for index, num := range nums {
		diff := target - num
		if _, ok := hmap[diff]; ok {
			return []int{index, hmap[diff]}
		}
		hmap[num] = index
	}
	return nil
}
