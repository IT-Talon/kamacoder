package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := levelOrder(&TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(l)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	order(root, 0)
	return res
}

func order(root *TreeNode, index int) {
	if root == nil {
		return
	}
	if len(res) > index {
		res[index] = append(res[index], root.Val)
	} else {
		res = append(res, []int{root.Val})
	}

	order(root.Left, index+1)
	order(root.Right, index+1)
}

func order2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	l := list.New()
	l.PushBack(root)
	res := make([][]int, 0)
	for l.Len() > 0 {
		var temp []int
		length := l.Len()
		for i := 0; i < length; i++ {
			n := l.Remove(l.Front()).(*TreeNode)
			temp = append(temp, n.Val)
			if n.Left != nil {
				l.PushBack(n.Left)
			}
			if n.Right != nil {
				l.PushBack(n.Right)
			}
		}
		res = append(res, temp)
	}
	return res

}
