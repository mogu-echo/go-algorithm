package problem0104

import (
	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	return max(left, right) + 1
}
