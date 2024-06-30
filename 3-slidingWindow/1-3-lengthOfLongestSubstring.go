package slidingWindow

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	left := 0
	maxLength := 0
	for right, char := range s {
		if index, ok := charMap[char]; ok {
			left = max(index+1, left)
		}
		charMap[char] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
