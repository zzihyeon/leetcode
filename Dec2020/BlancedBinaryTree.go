package Dec2020

import "log"

func IsBalanced(root *TreeNode) bool {
	defer func() {
		if s := recover(); s != nil {
			log.Println(s)
		}
	}()
	leftMax := 0
	rightMax := 0
	ptr := root
	if ptr == nil {
		return true
	}
	dfsSearchHeight(ptr.Left, 1, &leftMax)
	dfsSearchHeight(ptr.Right, 1, &rightMax)
	if leftMax-rightMax < -1 || leftMax-rightMax > 1 {
		return false
	}
	flag := IsBalanced(ptr.Left)
	if flag == false {
		return false
	}
	flag = IsBalanced(ptr.Right)
	if flag == false {
		return false
	}
	return true
}

func dfsSearchHeight(ptr *TreeNode, height int, max *int) {
	defer func() {
		if s := recover(); s != nil {
			log.Println(s)
		}
	}()

	if ptr == nil {
		return
	}

	if ptr.Left == nil && ptr.Right == nil {
		if *max < height {
			*max = height
		}
		return
	}
	dfsSearchHeight(ptr.Left, height+1, max)
	dfsSearchHeight(ptr.Right, height+1, max)
}
