package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input struct {
		s string
		t string
	}

	output string
)

func Test_MinimumWindowSubstring(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			out: "BANC",
		},
		{
			in: input{
				s: "a",
				t: "a",
			},
			out: "a",
		},
		{
			in: input{
				s: "a",
				t: "aa",
			},
			out: "",
		},
		{
			in: input{
				s: "ab",
				t: "b",
			},
			out: "b",
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			assert.Equal(t, string(test.out), minWindow(test.in.s, test.in.t))
		})
	}
}
