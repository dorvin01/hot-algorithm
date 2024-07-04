package subsequence

func maxSlidingWindow(nums []int, k int) []int {
	var stack []int
	push := func(index int) {
		for len(stack) > 0 && nums[index] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans := []int{nums[stack[0]]}
	for i := k; i < len(nums); i++ {
		push(i)
		for stack[0] <= i-k {
			stack = stack[1:]
		}
		ans = append(ans, nums[stack[0]])
	}
	return ans
}
