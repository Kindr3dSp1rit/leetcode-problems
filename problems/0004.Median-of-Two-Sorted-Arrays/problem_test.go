package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	input struct {
		nums1 []int
		nums2 []int
	}

	output float64
)

func Test_MedianOfTwoSortedArrays(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			out: 2,
		},
		{
			in: input{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			out: 2.5,
		},
		{
			in: input{
				nums1: []int{0, 0, 0, 0, 0},
				nums2: []int{-1, 0, 0, 0, 0, 0, 1},
			},
			out: 0,
		},
		{
			in: input{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			out: 2.5,
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, float64(test.out), findMedianSortedArrays(test.in.nums1, test.in.nums2))
		})
	}
}
