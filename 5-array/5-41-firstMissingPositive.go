package array

func firstMissingPositive(nums []int) int {
	oneExist := false
	for index := range nums {
		if nums[index] == 1 {
			oneExist = true
		}
		if nums[index] <= 0 || nums[index] > len(nums) {
			nums[index] = 1
		}
	}
	if !oneExist {
		return 1
	}
	for index := range nums {
		position := max(nums[index], -nums[index]) - 1
		nums[position] = -max(nums[position], -nums[position])
	}
	for index := range nums {
		if nums[index] > 0 {
			return index + 1
		}
	}
	return len(nums) + 1
}
