package leetcode

var solved = make(map[string]int)

func numDecodings(s string) int {
	res := 0

	if v, ok := solved[s]; ok {
		return v
	}

	if len(s) == 0 {
		return 1
	}
	if s[0] == byte('0') {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	res += numDecodings(s[1:])

	if s[0] == byte('1') || (s[0] == byte('2') && s[1] <= byte('6')) {
		res += numDecodings(s[2:])
	}

	solved[s] = res

	return res
}
