package leetcode

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}

	for _, v := range strs {
		key := getAnagramKey(v)
		groups[key] = append(groups[key], v)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}

	return res
}

// input strings only have lowercase english letters
func getAnagramKey(s string) string {
	letterCount := make(map[rune]int)
	resLen := 0
	for _, l := range s {
		letterCount[l]++
		resLen++
	}
	res := make([]rune, 0, resLen)
	for letter := 'a'; letter <= 'z'; letter++ {
		count := letterCount[letter]
		for i := 0; i < count; i++ {
			res = append(res, letter)
		}
	}

	return string(res)
}
