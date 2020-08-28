package problem1022

import "github.com/mogu-echo/go-algorithm/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, b int) int {
	if node == nil {
		return 0
	}
	b = b*2 + node.Val
	if node.Left == node.Right { // both nil
		return b
	}
	return dfs(node.Left, b) + dfs(node.Right, b)
}
