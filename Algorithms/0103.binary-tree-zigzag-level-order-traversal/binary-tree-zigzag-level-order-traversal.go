package problem0103

import (
	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		// 与 102 题相比，level 的奇偶不同，append 会不同。
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int, len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		if toggle {
			reverse(list)
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
