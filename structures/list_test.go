package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LinkedList2Slice(t *testing.T) {
	testCases := map[string]struct {
		in  *ListNode[int]
		out []int
	}{
		"nil input": {
			in:  nil,
			out: nil,
		},
		"normal input": {
			in: &ListNode[int]{
				Val: 1,
				Next: &ListNode[int]{
					Val: 2,
					Next: &ListNode[int]{
						Val: 3,
					},
				},
			},
			out: []int{1, 2, 3},
		},
		"panic on threshold": {
			in:  nil,
			out: make([]int, loopDetectionThreshold+1),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if len(test.out) > loopDetectionThreshold {
				assert.Panics(t, func() { LinkedList2Slice[int](Slice2LinkedList[int](test.out)) })
			} else {
				assert.Equal(t, test.out, LinkedList2Slice[int](test.in))
			}
		})
	}
}
