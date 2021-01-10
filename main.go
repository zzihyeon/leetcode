package main

import (
	"log"

	"main.go/Dec2020"
	"main.go/Dec2021"
	"main.go/types"
)

func main() {
	// valData := []int{9, 5, 4, 5, 0, 2, 6, 2, 5, 0, 8, 3, 9, 2, 3, 1, 1, 0, 4, 5, 4, 2, 2, 6, 4, 0, 0, 1, 7, 0, 5, 4, 7, 0, 0, 7, 0, 1, 5, 6, 1, 0, 0, 0, 0, 9, 2, 0, 9, 7, 2, 1, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 3, 0, 0, 0, 8, 0, 1, 0, 0, 8, 0, 0, 0, 0, 2, 0, 8, 7}
	valData := []int{1, 1, 2}
	// data := types.TreeNode{}
	data := types.ListNode{}
	// Dec2020.MakeNodeData(&data, valData)
	Dec2020.MakeListNodeData(&data, valData)

	// result := Dec2020.IsBalanced(&data)
	// result := Dec2020.PseudoPalindromicPaths(&data)
	// result := Dec2020.NextGreaterElement(1999999999)
	// result := Dec2020.SwapPairs(&data)
	// data := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// Dec2020.GameOfLife(data)
	// result := Dec2020.LargestRectangleArea(valData)
	// result := Dec2021.CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}})
	// result := Dec2021.CanFormArray([]int{1, 3, 5, 7}, [][]int{{2, 4, 6, 8}})
	// result := Dec2021.CountArrangement(6)
	// Dec2021.MergeTwoSortedLists()
	result := Dec2021.DeleteDuplicates(&data)
	log.Println(result)
}
