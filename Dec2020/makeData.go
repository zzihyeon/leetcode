package Dec2020

import "main.go/types"

func MakeNodeData(root *types.TreeNode, val []int) {
	queue := []*types.TreeNode{}
	queueCnt := 0
	for i, ele := range val {
		if i == 0 {
			root.Val = ele
			queue = append(queue, root)
		} else {
			var newNode *types.TreeNode
			if ele == 0 {
				newNode = nil
			} else {
				newNode = &types.TreeNode{Val: ele}
				queue = append(queue, newNode)
			}
			if queueCnt == 2 {
				queue = queue[1:]
				queueCnt = 0
			}

			parent := queue[0]
			if queueCnt == 0 {
				parent.Left = newNode
			} else if queueCnt == 1 {
				parent.Right = newNode
			}

			queueCnt++
		}
	}
}

func MakeListNodeData(root *types.ListNode, val []int) {
	ptr := &types.ListNode{}
	for i, ele := range val {
		newNode := &types.ListNode{}
		newNode.Val = ele
		if i == 0 {
			root.Val = ele
			ptr = root
		} else {
			ptr.Next = newNode
			ptr = ptr.Next
		}
	}
}
