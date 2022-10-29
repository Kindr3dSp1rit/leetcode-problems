package leetcode

func minWindow(s string, t string) string {
	// keep track of number of required characters and
	// total number of characters to fill
	chars := make(map[byte]int)
	sign := 0
	for i := 0; i < len(t); i++ {
		c := t[i]
		if _, ok := chars[c]; !ok {
			chars[c] = 0
		}
		chars[c]++
		sign++
	}

	w1 := 0     // beginning of window
	w2 := 0     // end of window
	minlen := 0 // minimal length of suitable window so far
	minw1 := 0  // first index of minimal window

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := chars[c]; ok {
			// if one of the required characters if encountered,
			// lower its count and signature number (if char is not filled)
			chars[c]--
			if chars[c] >= 0 {
				sign--
				// if requirements for the encountered character are already filled
				// try to cut the current window from the front
			}
			if chars[c] <= 0 {
				for {
					// cut window from the front until encountering character that is not yet filled
					if v, ok := chars[s[w1]]; ok {
						if v >= 0 {
							break
						} else {
							chars[s[w1]]++
						}
					}
					w1++
				}
			}
		}
		if sign == 0 && (w2-w1+1 < minlen || minlen == 0) {
			minlen = w2 - w1 + 1
			minw1 = w1
		}
		w2++
	}
	if minlen == 0 {
		return ""
	}
	return s[minw1 : minw1+minlen]
}
