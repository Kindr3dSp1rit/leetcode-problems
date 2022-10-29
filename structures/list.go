package structures

import "fmt"

const (
	loopDetectionThreshold = 1000
)

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func LinkedList2Slice[T any](head *ListNode[T]) []T {
	length := 0

	var res []T
	for head != nil {
		length++
		if length > loopDetectionThreshold {
			panic(fmt.Sprintf("result list too long, possible loop: %v\n", head))
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func Slice2LinkedList[T any](nums []T) *ListNode[T] {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode[T]{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode[T]{Val: v}
		t = t.Next
	}
	return l.Next
}
