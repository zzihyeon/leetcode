package Dec2020

func SwapPairs(head *ListNode) *ListNode {
	prev := head
	for ptr := head; ptr != nil && ptr.Next != nil; {
		if prev != head {
			prev.Next = ptr.Next
		} else {
			head = ptr.Next
		}
		post := ptr.Next
		ptr.Next = post.Next
		post.Next = ptr
		prev = ptr
		ptr = ptr.Next
	}
	return head
}
