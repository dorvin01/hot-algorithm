package doublePointer

func moveZeroes(nums []int) {
	left := 0
	numsCount := len(nums)
	for right := 0; right < numsCount; right++ {
		if nums[right] == 0 {
			continue
		}
		if left < right {
			nums[left] = nums[right]
		}
		left++
	}
	for i := left; i < numsCount; i++ {
		nums[i] = 0
	}
}
