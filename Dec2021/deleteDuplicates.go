package Dec2021

import (
	"main.go/types"
)

func DeleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newNode *types.ListNode
	var newNodePtr *types.ListNode
	prevVal := 101
	ptr := head
	for ; ptr.Next != nil; ptr = ptr.Next {
		if (ptr.Val != ptr.Next.Val) && (ptr.Val != prevVal) {
			if newNode == nil {
				newNode = ptr
				newNodePtr = ptr
			} else {
				newNodePtr.Next = ptr
				newNodePtr = newNodePtr.Next
			}
		}
		prevVal = ptr.Val
	}
	if newNode != nil {
		newNodePtr.Next = nil
		if prevVal != ptr.Val {
			newNodePtr.Next = ptr
		}
	} else if prevVal != ptr.Val {
		newNode = ptr
	}
	return newNode
}
