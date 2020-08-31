package problem0124

import (
	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	var dfs func(*TreeNode) int
	// 返回，以 root 为起点，所有可能路径的 sum 值中的最大值。
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val + right
		if maxSum < sum {
			maxSum = sum
		}

		return max(left, right) + root.Val
	}

	dfs(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ResultType struct {
	SinglePath, MaxPath int
}

func maxPathSum2(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			0, -(1 << 31),
		}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	result := ResultType{}
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}
