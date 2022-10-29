package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	input []string

	output [][]string
)

func Test_GroupAnagrams(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			out: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			in:  []string{""},
			out: [][]string{{""}},
		},
		{
			in:  []string{"a"},
			out: [][]string{{"a"}},
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, [][]string(test.out), groupAnagrams(test.in))
		})
	}
}
