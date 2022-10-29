package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input struct {
		plantTime []int
		growTime  []int
	}
	output int
)

func Test_EarliestFullBloom(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				plantTime: []int{1, 4, 3},
				growTime:  []int{2, 3, 1},
			},
			out: 9,
		},
		{
			in: input{
				plantTime: []int{1, 2, 3, 2},
				growTime:  []int{2, 1, 2, 1},
			},
			out: 9,
		},
		{
			in: input{
				plantTime: []int{1},
				growTime:  []int{1},
			},
			out: 2,
		},
		{
			in: input{
				plantTime: []int{3, 11, 29, 4, 4, 26, 26, 12, 13, 10, 30, 19, 27, 2, 10},
				growTime:  []int{10, 13, 22, 17, 18, 15, 21, 11, 24, 14, 18, 23, 1, 30, 6},
			},
			out: 227,
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			assert.Equal(t, int(test.out), earliestFullBloom(test.in.plantTime, test.in.growTime))
		})
	}
}
