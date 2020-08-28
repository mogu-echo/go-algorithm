package problem0100

import (
	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q== nil {
		return true 
	}

	if p == nil || q == nil {
		return false 
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) &&  isSameTree(p.Right, q.Right)
}
