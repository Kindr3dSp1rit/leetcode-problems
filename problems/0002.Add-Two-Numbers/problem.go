package leetcode

import (
	"leetcode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode[int]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	currentNode := head
	prevNode := head
	carry := 0

	currentL1Node := l1
	currentL2Node := l2
	for currentL1Node != nil && currentL2Node != nil {
		sum := currentL1Node.Val + currentL2Node.Val + carry
		carry = sum / 10
		currentNode.Val = sum % 10
		currentL1Node = currentL1Node.Next
		currentL2Node = currentL2Node.Next
		prevNode = currentNode
		currentNode = new(ListNode)
		prevNode.Next = currentNode
	}
	if currentL1Node == nil {
		currentL1Node = currentL2Node
	}
	for currentL1Node != nil {
		sum := currentL1Node.Val + carry
		carry = sum / 10
		currentNode.Val = sum % 10
		currentL1Node = currentL1Node.Next
		prevNode = currentNode
		currentNode = new(ListNode)
		prevNode.Next = currentNode
	}
	if carry > 0 {
		currentNode.Val += carry
	}
	if currentNode.Val == 0 {
		prevNode.Next = nil
	}
	return head
}
