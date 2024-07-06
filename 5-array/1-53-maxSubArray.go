package array

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	var sum int
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		res = max(res, sum)
	}
	return res
}
