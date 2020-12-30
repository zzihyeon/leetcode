package main

import (
	"log"

	"main.go/Dec2020"
)

func main() {
	// valData := []int{9, 5, 4, 5, 0, 2, 6, 2, 5, 0, 8, 3, 9, 2, 3, 1, 1, 0, 4, 5, 4, 2, 2, 6, 4, 0, 0, 1, 7, 0, 5, 4, 7, 0, 0, 7, 0, 1, 5, 6, 1, 0, 0, 0, 0, 9, 2, 0, 9, 7, 2, 1, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 3, 0, 0, 0, 8, 0, 1, 0, 0, 8, 0, 0, 0, 0, 2, 0, 8, 7}
	// data := Dec2020.TreeNode{}
	// Dec2020.MakeNodeData(&data, valData)

	// result := Dec2020.IsBalanced(&data)
	// result := Dec2020.PseudoPalindromicPaths(&data)
	result := Dec2020.NextGreaterElement(1999999999)
	log.Println(result)
}
