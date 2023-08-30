package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var preorder, inorder string
		_, err := fmt.Scan(&preorder, &inorder)
		if err != nil {
			return
		}
		tree := buildTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
		fmt.Println(postorderTraversal(tree))
	}
}

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

// 根据前序和中序遍历结果构造二叉树
func buildTree(preorder, inorder string, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	//根节点值为前序遍历结果的第一个字符
	root := &TreeNode{Val: string(preorder[preStart])}

	// 在中序遍历结果中找到根节点的位置
	var rootIndex int
	for rootIndex = inStart; rootIndex <= inEnd; rootIndex++ {
		if inorder[rootIndex] == preorder[preStart] {
			break
		}
	}
	// 递归构造左子树和右子树
	leftTreeSize := rootIndex - inStart
	root.Left = buildTree(preorder, inorder, preStart+1, preStart+leftTreeSize, inStart, rootIndex-1)
	root.Right = buildTree(preorder, inorder, preStart+leftTreeSize+1, preEnd, rootIndex+1, inEnd)

	return root
}

func postorderTraversal(root *TreeNode) string {
	var res []string
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return strings.Join(res, "")
}
