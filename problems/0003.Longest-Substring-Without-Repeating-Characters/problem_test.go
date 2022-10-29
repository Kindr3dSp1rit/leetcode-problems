package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input  string
	output int
)

func Test_LongestSubstringWithoutRepeatingCharacters(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in:  "abcabcbb",
			out: 3,
		},
		{
			in:  "bbbbb",
			out: 1,
		},
		{
			in:  "pwwkew",
			out: 3,
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, int(test.out), lengthOfLongestSubstring(string(test.in)))
		})
	}
}
