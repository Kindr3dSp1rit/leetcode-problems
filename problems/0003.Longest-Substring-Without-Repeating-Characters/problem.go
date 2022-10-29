package leetcode

func lengthOfLongestSubstring(s string) int {
	subsequence := make(map[byte]int, 0)
	maxLen := 0
	currentLen := 0
	i := 0
	for i < len(s) {
		if idx, ok := subsequence[s[i]]; ok {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			subsequence = make(map[byte]int, 0)
			i = idx + 1
			if i >= len(s) {
				return maxLen
			}
			currentLen = 0
		}
		subsequence[s[i]] = i
		currentLen += 1
		i += 1
	}
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}
