package Dec2021

import "main.go/types"

func MergeTwoLists(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	l1Ptr := l1
	l2Ptr := l2
	var root *types.ListNode
	if l1Ptr == nil {
		return l2Ptr
	}
	if l2Ptr == nil {
		return l1Ptr
	}
	if l1Ptr.Val < l2Ptr.Val {
		root = l1Ptr
		l1Ptr = l1Ptr.Next
	} else {
		root = l2Ptr
		l2Ptr = l2Ptr.Next
	}
	for ptr := root; ; ptr = ptr.Next {
		if l1Ptr == nil {
			ptr.Next = l2Ptr
			return root
		}
		if l2Ptr == nil {
			ptr.Next = l1Ptr
			return root
		}
		if l1Ptr.Val < l2Ptr.Val {
			ptr.Next = l1Ptr
			l1Ptr = l1Ptr.Next
		} else {
			ptr.Next = l2Ptr
			l2Ptr = l2Ptr.Next
		}
	}
}
