package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input struct {
		img1 [][]int
		img2 [][]int
	}

	output int
)

func Test_ImageOverlap(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				img1: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
				img2: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			},
			out: 3,
		},
		{
			in: input{
				img1: [][]int{{1}},
				img2: [][]int{{1}},
			},
			out: 1,
		},
		{
			in: input{
				img1: [][]int{{0}},
				img2: [][]int{{0}},
			},
			out: 0,
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, int(test.out), largestOverlap(test.in.img1, test.in.img2))
		})
	}
}
