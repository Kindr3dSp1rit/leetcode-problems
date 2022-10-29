package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input struct {
		nums   []int
		target int
	}

	output []int
)

func Test_TwoSum(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			out: output([]int{0, 1}),
		},
		{
			in: input{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			out: output([]int{1, 2}),
		},
		{
			in: input{
				nums:   []int{3, 3},
				target: 6,
			},
			out: output([]int{0, 1}),
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			res := twoSum(test.in.nums, test.in.target)
			assert.Equal(t, test.out, output(res))
		})
	}
}
