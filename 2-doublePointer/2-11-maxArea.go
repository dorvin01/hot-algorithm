package doublePointer

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res, area int
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
