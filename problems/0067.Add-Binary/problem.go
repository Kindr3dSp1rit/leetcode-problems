package leetcode

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	big := a
	small := b

	if len(a) < len(b) {
		big = b
		small = a
	}

	for i := len(big) - len(small); i > 0; i-- {
		small = "0" + small
	}
	var result string
	var haveOne int
	for j := len(big); j > 0; j-- {
		sum := parseInt(big[j-1:j]) + parseInt(small[j-1:j]) + haveOne
		haveOne = 0
		if sum >= 2 {
			haveOne = 1
			sum = sum % 2
		}

		result = fmt.Sprintf("%d%s", sum, result)
	}

	if haveOne == 1 {
		result = fmt.Sprintf("%d%s", haveOne, result)
	}

	return result
}

func parseInt(target string) int {
	result, _ := strconv.Atoi(target)
	return result
}
