package Dec2020

import "main.go/types"

/**
 * Definition for a binary tree node.
 * type types.TreeNode struct {
 *     Val int
 *     Left *types.TreeNode
 *     Right *types.TreeNode
 * }
 */ //

func PseudoPalindromicPaths(root *types.TreeNode) int {
	cnt := 0
	makePermutation(root, []int{}, &cnt)
	return cnt
}

func makePermutation(ptr *types.TreeNode, permuArray []int, cnt *int) {
	if ptr == nil {
		return
	}
	nextArray := append(permuArray, ptr.Val)
	if ptr.Left == nil && ptr.Right == nil {
		if res := checkPalindromic(nextArray); res == true {
			*cnt = *cnt + 1
		}
	}
	makePermutation(ptr.Left, nextArray, cnt)
	makePermutation(ptr.Right, nextArray, cnt)
}

func checkPalindromic(permuArray []int) bool {
	checkMap := make(map[int]int)
	for _, ele := range permuArray {
		checkMap[ele]++
	}
	oddCount := 0
	for _, ele := range checkMap {
		if ele%2 == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}
