package template

import (
	"fmt"

	"github.com/mogu-echo/go-algorithm/kit"
)

type TreeNode = kit.TreeNode

//前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

//前序非递归
func preorderTraversalIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//中序非递归
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = root.Right
	}
	return result
}

//后序非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

//DFS 深度搜索-从上到下
func upDownTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

//DFS 深度搜索-从下向上（分治法）
func downUpTraversal(root *TreeNode) []int {
	return divideAndConquer(root)
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

//BFS 层次遍历
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

// 分治法应用
// 先分别处理局部，再合并结果

// 适用场景
// 	快速排序
// 	归并排序
// 	二叉树相关问题

// 分治法模板
// 	递归返回条件
// 	分段处理
// 	合并结果

// func traversal(root *TreeNode) T{
// 	// nil or leaf
// 	if root == nil {
//         // do something and return
// 	}

// 	// Divide
// 	T left = traversal(root.Left)
// 	T right = traversal(root.Right)

// 	// Conquer
// 	T result = Merge from left and right

// 	return result
// }

// 通过分治法遍历二叉树
// func downUpTraversal(root *TreeNode) []int {
// 	return divideAndConquer(root)
// }

// func divideAndConquer(root *TreeNode) []int {
// 	result := make([]int, 0)
// 	if root == nil {
// 		return result
// 	}
// 	left := divideAndConquer(root.Left)
// 	right := divideAndConquer(root.Right)
// 	result = append(result, root.Val)
// 	result = append(result, left...)
// 	result = append(result, right...)
// 	return result
// }

// 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) (result []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

// 快速排序
func QuickSort(nums []int) []int {
	// 思路：把一个数组分为左右两段，左段小于右段，类似分治法没有合并过程
	quickSort(nums, 0, len(nums)-1)
	return nums

}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法：divide
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, i, end)
	return i
}
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	// 返回条件处理
	if root == nil {
		return 0
	}
	// divide：分左右子树分别计算
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// conquer：合并左右子树结果
	if left > right {
		return left + 1
	}
	return right + 1
}

// https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}
func maxDepth(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
type ResultType struct {
	SinglePath int // 保存单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}
func helper(root *TreeNode) ResultType {
	// check
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}
	// Divide
	left := helper(root.Left)
	right := helper(root.Right)

	// Conquer
	result := ResultType{}
	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	// 求两边加根最大值
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// func levelOrder(root *TreeNode) [][]int {
// 	result := make([][]int, 0)
// 	if root == nil {
// 		return result
// 	}
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		list := make([]int, 0)
// 		// 为什么要取length？
// 		// 记录当前层有多少元素（遍历当前层，再添加下一层）
// 		l := len(queue)
// 		for i := 0; i < l; i++ {
// 			// 出队列
// 			level := queue[0]
// 			queue = queue[1:]
// 			list = append(list, level.Val)
// 			if level.Left != nil {
// 				queue = append(queue, level.Left)
// 			}
// 			if level.Right != nil {
// 				queue = append(queue, level.Right)
// 			}
// 		}
// 		result = append(result, list)
// 	}
// 	return result
// }
