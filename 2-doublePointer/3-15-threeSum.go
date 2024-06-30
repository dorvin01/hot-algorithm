package doublePointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	numsCount := len(nums)
	var res [][]int
	for i := 0; i < numsCount; i++ {
		if (i > 0) && (nums[i] == nums[i-1]) {
			continue
		}
		left, right := i+1, numsCount-1
		target := -nums[i]
		for ; left < right; left++ {
			if (left != i+1) && (nums[left] == nums[left-1]) {
				continue
			}
			for (nums[left]+nums[right] > target) && (left < right) {
				right--
			}
			if left == right {
				break
			}
			if target == nums[left]+nums[right] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
			}
			left++
		}
	}
	return res
}
