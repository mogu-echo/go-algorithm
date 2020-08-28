package problem0530

import (
	"github.com/mogu-echo/go-algorithm/kit"
)
type TreeNode = kit.TreeNode

type state struct{
    minDiff , previous int
}

func getMinimumDifference(root *TreeNode) int {
    st:=state{1024,1024}
    search(root, &st)
    return st.minDiff
}

// NOTICE: BST 的递归遍历方法
// 特别好
func search(root *TreeNode, st *state){
    if root == nil{
        return
	}

    search(root.Left, st)
    
    newDiff:=diff(st.previous, root.Val)
    if st.minDiff> newDiff{
        st.minDiff = newDiff
    }
    
	st.previous = root.Val

    search(root.Right, st)
}

func diff(a, b int) int{
    if a > b{
        return a - b
    }
	return b - a
}
