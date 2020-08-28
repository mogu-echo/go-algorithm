package problem1038

import "github.com/mogu-echo/go-algorithm/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

func bstToGst(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	node.Val += dfs(node.Right, sum)
	return dfs(node.Left, node.Val)
}
