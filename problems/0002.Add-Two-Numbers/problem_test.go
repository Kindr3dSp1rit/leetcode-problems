package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/structures"
)

type (
	input struct {
		l1 []int
		l2 []int
	}

	output []int
)

func Test_AddTwoNumbers(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			out: output([]int{7, 0, 8}),
		},
		{
			in: input{
				l1: []int{0},
				l2: []int{0},
			},
			out: output([]int{0}),
		},
		{
			in: input{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			out: output([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			l1, l2 := structures.Slice2LinkedList[int](test.in.l1), structures.Slice2LinkedList[int](test.in.l2)
			res := addTwoNumbers(l1, l2)
			assert.Equal(t, test.out, output(structures.LinkedList2Slice(res)))
		})
	}
}
