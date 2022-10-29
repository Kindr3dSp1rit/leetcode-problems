package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	input  string
	output int
)

func Test_DecodeWays(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in:  "12",
			out: 2,
		},
		{
			in:  "226",
			out: 3,
		},
		{
			in:  "06",
			out: 0,
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			solved = make(map[string]int)
			assert.Equal(t, int(test.out), numDecodings(string(test.in)))
		})
	}
}
