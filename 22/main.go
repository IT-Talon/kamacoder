package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	nodeMap = make(map[string]NodeInfo)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		l := strings.Fields(string(line))
		nodes = append(nodes, l[0])
		nodeMap[l[0]] = NodeInfo{
			val:   l[0],
			left:  l[1],
			right: l[2],
		}
	}
	tree := buildTree(nodes[0])
	fmt.Println(strings.Join(preorderTraversal(tree), ""))
	fmt.Println(strings.Join(inorderTraversal(tree), ""))
	fmt.Println(strings.Join(postorderTraversal(tree), ""))

}

var nodeMap map[string]NodeInfo
var nodes []string

type NodeInfo struct {
	val   string
	left  string
	right string
}

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

// 根据输入数据构造二叉树
func buildTree(val string) *TreeNode {
	root := &TreeNode{Val: val}
	left, _ := strconv.Atoi(nodeMap[val].left)
	right, _ := strconv.Atoi(nodeMap[val].right)
	if left == 0 {
		root.Left = nil
	} else {
		root.Left = buildTree(nodes[left-1])
	}
	if right == 0 {
		root.Right = nil
	} else {
		root.Right = buildTree(nodes[right-1])
	}

	return root
}

// 前序遍历
func preorderTraversal(root *TreeNode) (vals []string) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// 中序遍历
func inorderTraversal(root *TreeNode) (res []string) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 后序遍历
func postorderTraversal(root *TreeNode) (res []string) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}
