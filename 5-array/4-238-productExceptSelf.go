package array

func productExceptSelf(nums []int) []int {
	leftProduct := []int{nums[0]}
	rightProduct := []int{0}
	res := []int{0}
	for index := 1; index < len(nums); index++ {
		leftProduct = append(leftProduct, leftProduct[index-1]*nums[index])
		// init
		rightProduct = append(rightProduct, 0)
		if index == len(nums)-1 {
			rightProduct[index] = nums[index]
		}
		res = append(res, 0)
	}
	for index := len(nums) - 2; index >= 0; index-- {
		rightProduct[index] = nums[index] * rightProduct[index+1]
	}
	res[0] = rightProduct[1]
	res[len(nums)-1] = leftProduct[len(nums)-2]
	for index := 1; index < len(nums)-1; index++ {
		res[index] = leftProduct[index-1] * rightProduct[index+1]
	}
	return res
}
