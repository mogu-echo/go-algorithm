package problem0098

import (
	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	// Go int 类型的最小值与最大值
	MIN, MAX := -1<<63, 1<<63-1

	return recur(MIN, MAX, root)
}

// 以递归的方式，检查 root.Val 是否在 (min, max) 范围内。
func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}

func isValidBST2(root *TreeNode) bool {
	result := inOrder(root)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, inOrder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inOrder(root.Right)...)
	return result
}
