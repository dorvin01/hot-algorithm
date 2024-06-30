package doublePointer

func trap(height []int) int {
	var res int
	stack := []int{}
	for right := range height {
		for (len(stack) > 0) && (height[right] > height[stack[len(stack)-1]]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := right - left - 1
			height := min(height[left], height[right]) - height[top]
			res += width * height
		}
		stack = append(stack, right)
	}
	return res
}
