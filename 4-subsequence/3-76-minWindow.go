package subsequence

import (
	"math"
)

func minWindow(s string, t string) string {
	var sCount [58]int
	targetCount := count(t)
	// right is set to -1 to ensure that the initial position can be added 1 in count
	left, right := 0, -1
	maxLeft, minRight := 0, math.MaxInt
	for left < len(s) {
		if isSameCount(sCount, targetCount) {
			if right-left < minRight-maxLeft {
				maxLeft = left
				minRight = right
			}
			sCount[int(s[left]-'A')] -= 1
			left++
		} else {
			if right == len(s)-1 {
				break
			}
			right++
			sCount[int(s[right]-'A')] += 1
		}
	}
	if minRight == math.MaxInt {
		return ""
	}
	return s[maxLeft : minRight+1]
}

func count(s string) [58]int {
	var res [58]int
	for index := range s {
		position := int(s[index] - 'A')
		res[position] += 1
	}
	return res
}

func isSameCount(sCount [58]int, targetCount [58]int) bool {
	for i := 0; i < 58; i++ {
		if sCount[i] < targetCount[i] {
			return false
		}
	}
	return true
}
